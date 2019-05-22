package main

import (
	"fmt"
	"math/rand"
	"time"
)

var tme int

type facil struct {
	queue       int //число сообщений, ожидающих передачи
	chan1active bool
	chan2active bool
	qt          []int
	qi          int
	repairs     []int
	breaks      []int
	bi          int
}

const (
	emtyque int = iota //очередь пустая передавать нечего
	mainchn            //основной канал
	broken
	sparchn //резервный канал
	updstat //проверка состояний
)

func getTotalIn(q []int, curtime int) int {
	t := 0
	for _, v := range q {
		if curtime > v {
			t++
		}
	}
	return t
}

func (fc facil) UpdateState(curstate int, curtime int) (int, int) {
	var state int
	switch curstate {
	case emtyque:
		if fc.chan1active {
			state = mainchn
		} else if fc.chan2active {
			state = sparchn
		}
	case mainchn:
		if curtime > fc.breaks[fc.bi] { //произошел разрыв
			fc.chan1active = false
			curtime = fc.breaks[fc.bi]
			state = sparchn
		} else if curtime < fc.qt[fc.qi] { //сообщение передано до прихода следующего
			curtime = fc.qt[fc.qi]
			fc.qi++
			if fc.queue > 0 {
				state = mainchn
			} else if fc.queue == 0 {
				state = emtyque
			}
		}
	case sparchn:
		if curtime < fc.qt[fc.qi] { //сообщение передано до прихода следующего
			curtime = fc.qt[fc.qi]
			fc.qi++
			if fc.queue > 0 {
				state = mainchn
			} else if fc.queue == 0 {
				state = emtyque
			}
		}
	}
}

func Generate(n, delta, fintime int) []int {
	var q []int
	for qt := n + rand.Intn(delta); qt < fintime; qt += n + rand.Intn(fintime) { //генерация поступающих сообщений
		if qt < fintime {
			q = append(q, qt)
		}
	}
	return q
}

func GenerateDelta(mas []int, n, delta int) []int {
	var ms []int
	for _, v := range mas {
		ms = append(ms, v+n+rand.Intn(delta))
	}
	return ms
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fintime := 500
	q := Generate(5, 9, fintime)
	fmt.Println(q, len(q))
	breaks := Generate(165, 71, fintime)
	fmt.Println(breaks, len(breaks))
	repairs := GenerateDelta(breaks, 24, 7)

	fc := facil{
		queue:       1,
		chan1active: true,
		chan2active: false,
		qt:          q,
		qi:          0,
		breaks:      breaks,
		bi:          0,
	}
	//State := mainchn
	//curt := 0
	nq := 0
	//nb := 0
	for State, curt := mainchn, 0; curt <= fintime; {
		//fmt.Println(fc.queue, State, curt)
		switch State {
		case emtyque:
			curt = q[nq]
			nq++
			if nq >= len(q) {
				break
			}
			fc.queue++
			State = fc.UpdateState(State, curt)
		case mainchn:
			curt += 4 + rand.Intn(7) // передача по 1-му каналу
			fc.UpdateState(State, curt)
		case sparchn:

		}
		//fmt.Println(fc.queue, State, curt)
	}
}

/*
switch State {
		case emtyque:
			fmt.Println("Waiting to message in", q[nq])
			curt = q[nq]
			if nq < len(q) {
				nq++
			} else {
				break
			}
			fc.queue++
			if fc.chan1active {
				State = mainchn
			} else if fc.chan2active {
				State = sparchn
			}
		case mainchn:
			curt += 4 + rand.Intn(7) // передача по 1-му каналу
			fmt.Println("Sending till", curt)
			if nq >= len(q) {
				break
			}
			if curt < q[nq] { //текущее время меньше чем время прихода след заявки
				fc.queue--
				if fc.queue == 0 {
					State = emtyque
				}
			} else {
				if nq < len(q) {
					nq++
				} else {
					break
				}
			}
		case sparchn:
			if !fc.chan2active {
				curt += 2 + rand.Intn(2)
				fc.chan2active = true
			}
			curt += 4 + rand.Intn(7) // передача по 2-му каналу
			fc.queue--
			if fc.queue == 0 {
				State = emtyque
			}
		}
		//fmt.Println(fc.queue, State, curt)
	}

*/
