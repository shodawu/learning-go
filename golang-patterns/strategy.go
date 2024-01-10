package main

import "fmt"

// 公司A規定員工上班以感應識別證方式簽到
// 公司B規定員工上班以把紙卡插入卡鐘方式簽到
// 每間公司都有「上班簽到」的相關規定

type Sensor struct {
}

func (s *Sensor) TouchIDCard() {
	fmt.Println("TouchIDCard")
}

type ComA struct {
	Com
	CheckInClock Sensor
}

func (ca *ComA) CheckIn() {
	fmt.Println("ComA CheckIn")
	ca.CheckInClock.TouchIDCard()
}

type PaperCard struct {
}

func (pc *PaperCard) InsertSalaryCard() {
	fmt.Println("InsertSalaryCard")
}

type ComB struct {
	Com
	CheckInEq PaperCard
}

func (cb *ComB) CheckIn() {
	fmt.Println("ComB CheckIn")
	cb.CheckInEq.InsertSalaryCard()
}

type ICheckIn interface {
	CheckIn()
}

type Com struct {
	ICheckIn
}

func main() {
	a := ComA{
		CheckInClock: Sensor{},
	}

	a.CheckIn()

	b := ComB{
		CheckInEq: PaperCard{},
	}

	b.CheckIn()
}
