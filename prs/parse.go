package prs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type tsk1 struct {
	Lambda   float64 `json:"Tc"`
	Mu       float64 `json:"Ts"`
	Waittime int     `json:"Tw"`
}

type tsk2 struct {
	Lambda float64 `json:"Ts"`
	Mu     float64 `json:"Tc"`
	Number int     `json:"n"`
}

type Quest struct {
	Task1 tsk1 `json:"task1"`
	Task2 tsk2 `json:"task2"`
}

func GetQuest() Quest {
	jsonFile, err := os.Open("task.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened task!")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var qst Quest

	json.Unmarshal(byteValue, &qst)

	qst.Task1.Lambda = 1. / qst.Task1.Lambda
	qst.Task1.Mu = 1. / qst.Task1.Mu

	return qst

}
