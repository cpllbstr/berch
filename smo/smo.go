package main

import (
	"fmt"
	"math/rand"
)

type facil struct {
	queue       int //число сообщений, ожидающих передачи
	chan1active bool
	chan2active bool
}

type Events struct {
	incmes int //взходящее сообщение
	breakc int //разрыв канала
	ch2act int //активация 2го канала
	ch1rep int //ремонт первого канала
}

const (
	emptyq int = iota //пустая очередь
	mesinq            //есть сообщение в очереди
	mainse            //отправление по главному каналу
	secnse            //отправление по второму каналу
	actsec            //активация второго канала
)

func NextEvent(starttime, n, delta int) int { //определяет время происхождения следующего события
	return starttime + n + rand.Intn(delta)
}

func GenerateFirstEvents() Events { //определяет время происхождения следующего события
	breakch := 165 + rand.Intn(71)
	return Events{
		incmes: 5 + rand.Intn(8),
		breakc: breakch,
		ch2act: breakch + 2 + rand.Intn(2),
		ch1rep: breakch + 24 + rand.Intn(7),
	}
}
func SendMessage(ctime int) int {
	return ctime + 4 + rand.Intn(7)
}

func (ev *Events) CheckQueue(q, t int) (int, int) {
	var State int
	if t < ev.incmes {
		if q == 0 {
			State = emptyq
		} else {
			State = mesinq
		}
	} else if t >= ev.incmes {
		for t >= ev.incmes { //сколько сообщений успело придти в накопитель до окончания передачи
			ev.incmes = NextEvent(ev.incmes, 5, 8)
			q++
		}
		State = mesinq
	}
	return State, q
}

func main() {
	fintime := 3600 //конечное время моделирования
	rand.Seed(1336)
	State := mainse
	ev := GenerateFirstEvents()
	fc := facil{
		queue:       1,
		chan1active: true,
		chan2active: false,
	}
	for curtime := 0; curtime < fintime; {
		if fc.queue < 0 {
			panic("negative queue")
		}
		switch State {
		case emptyq:
			fmt.Println("Empty queue in", curtime)
			curtime = ev.incmes
			fc.queue++
			ev.incmes = NextEvent(ev.incmes, 5, 8)
			State = mesinq
		case mesinq:
			if fc.chan1active {
				State = mainse
			} else if fc.chan2active {
				State = secnse
			} else {
				State = actsec
			}
		case mainse:
			fmt.Println("Sending in main", curtime, "queue:", fc.queue, "|next message at ", ev.incmes)
			curtime = SendMessage(curtime)
			fmt.Println("->", curtime)
			if curtime > ev.breakc {
				curtime = ev.breakc
				fmt.Println("Break at", curtime)
				State = mesinq
				fc.chan1active = false
				ev.breakc = NextEvent(ev.breakc, 165, 71)
				continue
			}
			fc.queue--
			State, fc.queue = ev.CheckQueue(fc.queue, curtime)
		case secnse:
			if !fc.chan2active {
				State = actsec
				continue
			}
			fmt.Println("Sending in secondary", curtime, "queue:", fc.queue, "|next message at", ev.incmes)
			curtime = SendMessage(curtime)
			fmt.Println("->", curtime)
			fc.queue--
			State, fc.queue = ev.CheckQueue(fc.queue, curtime)
			if curtime > ev.ch1rep {
				fmt.Println("Chan1 repaired at", curtime)
				//State = mesinq
				fc.chan1active = true
				fc.chan2active = false
				ev.ch1rep = NextEvent(ev.breakc, 24, 7)
				continue
			}
		case actsec:
			curtime = ev.ch2act
			ev.ch2act = NextEvent(ev.breakc, 2, 2)
			fmt.Println("Second chan activated at", curtime)
			fc.chan2active = true
			State = secnse
		default:
			panic("Unknown state!")
		}
	}
}
