package main

import (
	"math"
)

func factorial(x float64) float64 {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func Index(ro float64, pow, den int) float64 {
	return math.Pow(ro, float64(pow)) / factorial(float64(den))
}

func CalcNav(P []float64, q int) float64 {
	N := 0.0
	p := len(P) - q - 1 //число операторов
	for i, Pi := range P {
		if i <= p {
			N += Pi * float64(i)
		} else {
			N += Pi * float64(p)
		}
	}
	return N
}

func CalcQav(P []float64, q int) float64 {
	Q := 0.0
	p := len(P) - q //число операторов
	for i, k := p, 1.; i < len(P); i++ {
		Q += k * P[i]
		k += 1.
	}
	return Q
}

//CalcRefusal - считает процент отказов при заданных параметра а так же коэфф-ты и P
/*
	ro - lambda/mu
	n - число операторов, либо число операторов + число мест в очереди
	q - число мест в очереди
	считать что функция работает корректно
*/
func CalcRefusal(ro float64, n, q int) (float64, []float64, []float64) {
	P0 := 1.0
	p := n - q
	Pind := []float64{1}
	//цикл работает верно
	for i := 1; i <= p; i++ {
		s := Index(ro, i, i)
		Pind = append(Pind, s)
		P0 += s
	}
	//этот вроде теперь правильно работает
	for i := 1; i <= q; i++ {
		s := Index(ro, p, p) * (math.Pow(ro/float64(p), float64(i)))
		Pind = append(Pind, s)
		P0 += s
	}
	P0 = 1. / P0
	P := []float64{P0}
	for i := 1; i < n+1; i++ {
		P = append(P, Pind[i]*P0)
	}
	return Index(ro, p, p) * math.Pow(ro/float64(p), float64(q)) * P[0], P, Pind //Pn, P - все P, Pind - индексы
}

//CalcP0InfQ - P0 для бесконечной очереди
func CalcP0InfQ(lmb, mu float64, n int) float64 {
	ro := lmb / mu
	P0 := 1.
	for i := 1; i <= n; i++ {
		s := Index(ro, i, i)
		P0 += s
	}
	P0 += Index(ro, n, n) * lmb / (float64(n)*mu - lmb)
	return 1. / P0
}

func CalcQavInf(lmb, mu, P0 float64, n int) float64 {
	a := lmb / (float64(n) * mu)
	return Index(lmb/mu, n, n) * P0 * (a / math.Pow(1-a, 2))
}

func CalcNavInf(lmb, mu, P0 float64, n int) float64 {
	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += Index(lmb/mu, i, i-1.)
	}
	return P0 * (sum + Index(lmb/mu, n, n-1.)*(lmb/(float64(n)*mu-lmb)))
}

func InfProd(lmb, mu, v float64, k, n int) float64 {
	prod := 1.
	for j := 1; j <= k; j++ {
		prod = prod * (float64(n)*mu + float64(j)*v)
	}
	return math.Pow(lmb, float64(k)) / prod
}

func CalcP0Time(lmb, mu, v float64, n, lim int) float64 {
	P0 := 0.0
	for i := 1; i <= n; i++ {
		P0 += Index(lmb/mu, i, i)
	}
	for k := 1; k <= lim; k++ {
		P0 += Index(lmb/mu, n, n) * InfProd(lmb, mu, v, k, n)
	}
	return 1. / P0
}

func CalcQTime(lmb, mu, P0, v float64, n, lim int) float64 {
	Q := 0.0
	for k := 1; k <= lim; k++ {
		Q += float64(k) * InfProd(lmb, mu, v, k, n) * Index(lmb/mu, n, n) * P0
	}
	return Q
}

func CalcNTime(lmb, mu, P0, v float64, n, lim int) float64 {
	N := 0.0
	for i := 1; i <= n; i++ {
		N += Index(lmb/mu, i, i-1)
	}
	for k := 1; k <= lim; k++ {
		N += Index(lmb/mu, n, n-1) * InfProd(lmb, mu, v, k, n)
	}
	return N * P0
}

//ЗАДАНИЕ 2

func Prod(i, n int) float64 {
	prod := 1
	for j := 1; j <= i; j++ {
		prod = prod * (n - j + 1)
	}
	return float64(prod)
}

// CalcP0z2 - m-число наладчиков, n-число станков
func CalcP0z2(lmb, mu float64, m, n int) float64 {
	P0 := 1.
	for i := 1; i <= m; i++ {
		P0 += Prod(i, n) * Index(lmb/mu, i, i)
	}
	for k := m + 1; k <= n; k++ {
		P0 += Prod(k, n) * Index(lmb/mu, k, m) / math.Pow(float64(m), float64(k-m))
	}
	return 1. / P0
}

func CalcNozh(lmb, mu, P0 float64, m, n int) float64 {
	N := 0.0
	for k := m + 1; k <= n; k++ {
		for i := 1; i <= n-m; i++ {
			N += float64(i) * Prod(k, n) * Index(lmb/mu, k, m) / math.Pow(float64(m), float64(k-m))
		}
	}
	return P0 * N
}

func CalcNprost(lmb, mu, P0 float64, m, n int) float64 {
	N := 0.0
	for i := 1; i <= m; i++ {
		N += float64(i) * Prod(i, n) * Index(lmb/mu, i, i)
	}
	for k := m + 1; k <= n; k++ {
		N += float64(k) * Prod(k, n) * Index(lmb/mu, k, m) / math.Pow(float64(m), float64(k-m))
	}
	return P0 * N
}

func CalcMzan(lmb, mu, P0 float64, m, n int) float64 {
	M := 0.0
	for i := 1; i <= m; i++ {
		M += float64(i) * Prod(i, n) * Index(lmb/mu, i, i)
	}
	for k := m + 1; k <= n; k++ {
		M += float64(m) * Prod(k, n) * Index(lmb/mu, k, m) / math.Pow(float64(m), float64(k-m))
	}
	return P0 * M
}
