package main

import (
	"fmt"
	"time"
)

type message interface {
	getMessage() string
}

func sendMsg(msg message) {
	fmt.Println(msg.getMessage())
}

type birthdayMsg struct {
	birthdayTime 	time.Time
	recipentName 	string
}

func (bm birthdayMsg) getMessage() string {
	return fmt.Sprintf("Happy Birthday %s! 🎉🎂", bm.recipentName)
}

type sendReport struct {
	reportName 		string
	numberofSends 	int
}

func (sr sendReport) getMessage() string {
	return fmt.Sprintf("Report %s has been sent %d times", sr.reportName, sr.numberofSends)
}

func test(m message) {
	sendMsg(m)
}

func main() {
	test(sendReport{
		reportName: "Sales Report",
		numberofSends: 5,
	})

	test(birthdayMsg{
		birthdayTime: time.Now(),
		recipentName: "John Doe",
	})
}