package main

import "fmt"

// 自助餐店提供的餐點有三類：主菜、配菜、主食
// 主菜類例如雞腿、排骨、魚等，每份50元。
// 配菜例如炒青菜、豆皮、蒸蛋等，每份15元。
// 主食有一份為半碗飯，每份10元。

type IMeal interface {
	Add(iPack int) int
}

type MainDish struct {
	IMeal
}

func (md *MainDish) Add(iPack int) int {
	return 50 * iPack
}

type SideDish struct {
	IMeal
}

func (sd *SideDish) Add(iPack int) int {
	return 15 * iPack
}

type Staple struct {
	IMeal
}

func (s *Staple) Add(iPack int) int {
	return 10 * iPack
}

type LucnchBox struct {
	price int
}

func (lb *LucnchBox) CalTotal() int {
	return lb.price
}

func (lb *LucnchBox) Order(i IMeal, iPack int) {
	lb.price += i.Add(iPack)
}
func main() {
	b := LucnchBox{}

	b.Order(&MainDish{}, 1)
	b.Order(&SideDish{}, 3)
	b.Order(&Staple{}, 2)
	fmt.Println(b.CalTotal())

}
