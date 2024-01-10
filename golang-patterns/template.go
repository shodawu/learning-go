package main

import "fmt"

// 便當店的客人會用"外帶,排骨飯"、"內用,雞腿飯"等方式點餐
// 依客人點餐包便當，便當包括餐點、餐具、飲料
// 營業流程是: 接收訂單-> 備餐 -> 交貨
// 依內用或外帶，於備餐時附的餐具與飲料會不一樣
// 內用的餐具是餐盤，飲料是湯
// 外帶的餐具是提袋，飲料是冷飲

type ILunchBox interface {
	Order(meal string)
	Prepare()
	Deliver()
}
type LunchBox struct {
	Meal  string
	Tool  string
	Drink string
	ILunchBox
}

func (lb *LunchBox) Order(meal string) {
	fmt.Println("接收訂單")
	lb.Meal = meal
}

type TakeOutBox struct {
	LunchBox
}

func (t *TakeOutBox) Prepare() {
	t.Tool = "提袋"
	t.Drink = "冷飲"
}

type StayBox struct {
	LunchBox
}

func (s *StayBox) Prepare() {
	s.Tool = "餐盤"
	s.Drink = "湯"
}

func (lb *LunchBox) Deliver() {
	fmt.Println("交貨", lb.Meal, lb.Tool, lb.Drink)
}

func main() {

	box := TakeOutBox{}

	box.Order("雞腿飯")
	fmt.Println("備餐訂單")
	box.Prepare()
	box.Deliver()

}
