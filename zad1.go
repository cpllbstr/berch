package main

import (
	"fmt"
	"os"

	"./prs"
)

func main() {

	tasks := prs.GetQuest()

	//Задание 1
	gnuplt, err1 := os.Create("./dat/refuses.dat")
	queue, err2 := os.Create("./dat/qref.dat")
	inf, err3 := os.Create("./dat/inf.dat")
	inftime, err4 := os.Create("./dat/inftime.dat")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		panic("Error opening file!")
	}
	defer gnuplt.Close()
	defer queue.Close()
	defer inf.Close()
	defer inftime.Close()

	n1prc := -1
	ro := float64(tasks.Task1.Lambda) / float64(tasks.Task1.Mu) //ro - приведенная интенсивность
	lmb := tasks.Task1.Lambda
	mu := tasks.Task1.Mu
	v := 1. / tasks.Task1.Waittime

	//Отказы без очереди
	for n := 1; ; n++ {
		ref, P, _ := CalcRefusal(ro, n, 0) //цикл перебора количества обслуживающих
		gnuplt.WriteString(fmt.Sprintf("%v, %v, %v\n", n, ref, CalcNav(P, 0)/float64(n)))
		if ref < 0.01 {
			n1prc = n
			break
		}
	}
	if n1prc < 1 { //число при котором доля отказов не превышает 1%
		panic("Error!")
	}

	//Отказы с очередью
	for q := 0; q <= n1prc; q++ {
		ref, P, _ := CalcRefusal(ro, n1prc, q) //цикл длины очереди
		//fmt.Println(len(P), n1prc)
		Qav := CalcQav(P, q)
		queue.WriteString(fmt.Sprintf("%v, %v, %v, %v, %v, %v, %v\n", q, ref, n1prc-q, CalcNav(P, q)/float64(n1prc-q), Qav, Qav/float64(q), Qav/tasks.Task1.Lambda))
	}

	//Бесконечная очередь, для сходимости число операторов 9>=n<=16
	nshod := 9 // n сходимости
	for i := nshod; i <= n1prc; i++ {
		P0 := CalcP0InfQ(lmb, mu, i)
		inf.WriteString(fmt.Sprintf("%v, %v, %v, %v\n", i, CalcQavInf(lmb, mu, P0, i), CalcNavInf(lmb, mu, P0, i)/float64(i), CalcQavInf(lmb, mu, P0, i)/lmb))
	}

	//Бесконечная очередь с временем ожидания
	lim := 50 //предел для расчета приблизительного произведения
	for i := 1; i <= n1prc; i++ {
		P0 := CalcP0Time(lmb, mu, v, i, lim)
		Q := CalcQTime(lmb, mu, P0, v, i, lim)
		N := CalcNTime(lmb, mu, P0, v, i, lim)
		inftime.WriteString(fmt.Sprintf("%v %v %v %v\n", i, N/float64(i), Q, Q/lmb))
	}

}
