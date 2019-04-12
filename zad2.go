package main

import (
	"fmt"
	"os"

	"./prs"
)

func main() {
	tasks := prs.GetQuest()
	//Задание 2
	lmb := tasks.Task2.Lambda
	mu := tasks.Task2.Mu
	n := tasks.Task2.Number

	file, err := os.Create("./dat/zad2.dat")
	if err != nil {
		panic("Error creating file!")
	}
	defer file.Close()

	for m := 1; m <= 5; m++ {
		P0 := CalcP0z2(lmb, mu, m, n)
		Nozh := CalcNozh(lmb, mu, P0, m, n)
		Nprost := CalcNprost(lmb, mu, P0, m, n)
		Mz := CalcMzan(lmb, mu, P0, m, n)
		file.WriteString(fmt.Sprintf("%v, %v, %v, %v, %v\n", m, Nozh/float64(n), Nprost/float64(n), Mz, Mz/float64(m)))
	}
}
