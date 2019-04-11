package main

import (
	"fmt"
	"math"
	"os"

	"./prs"
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

func main() {

	tasks := prs.GetQuest()

	//Задание 1
	gnuplt, err1 := os.Create("./dat/refuses.dat")
	queue, err2 := os.Create("./dat/qref.dat")
	if err1 != nil || err2 != nil {
		panic(err1)
	}
	defer gnuplt.Close()
	defer queue.Close()
	n1prc := -1
	ro := float64(tasks.Task1.Lambda) / float64(tasks.Task1.Mu) //ro - приведенная интенсивность

	//Отказы без очереди
	for n := 1; ; n++ {
		ref, P, _ := CalcRefusal(ro, n, 0) //цикл перебора количества обслуживающих
		gnuplt.WriteString(fmt.Sprintf("%v, %v, %v\n", n, ref, CalcNav(P, 0)/float64(n)))
		if ref < 0.01 {
			n1prc = n
			break
		}
	}
	fmt.Println("N1prc:", n1prc)
	if n1prc < 1 {
		panic("Error!")
	}
	//Отказы с очередью
	for q := 0; q <= n1prc; q++ {
		ref, P, _ := CalcRefusal(ro, n1prc, q) //цикл длины очереди
		//fmt.Println(len(P), n1prc)
		Qav := CalcQav(P, q)
		queue.WriteString(fmt.Sprintf("%v, %v, %v, %v, %v, %v, %v\n", q, ref, n1prc-q, CalcNav(P, q)/float64(n1prc-q), Qav, Qav/float64(q), Qav/tasks.Task1.Lambda))
	}
}
