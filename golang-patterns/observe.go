package main

import "fmt"

// 旅客想知道自己想搭的車現在在哪一站
// 只要公車所在站改變，就更新訊息給所有旅客、站牌
// 公車 <-> 旅客
// 公車 <-> 站牌

type Bus struct {
	Obs []IObserve
}

func (b *Bus) Register(o IObserve) {
	b.Obs = append(b.Obs, o)
}

func (b *Bus) Remove(o IObserve) {
	for i := range b.Obs {
		if b.Obs[i] == o {
			b.Obs = append(b.Obs[:i], b.Obs[i+1:]...)
			return
		}
	}
}

func (b *Bus) Notify(s string) {
	for _, o := range b.Obs {
		o.Update(s)
	}
}

type IObserve interface {
	Update(s string)
}

type Station struct {
}

func (st *Station) Update(s string) {
	fmt.Println("Station Update", s)
}

type Passanger struct {
}

func (p *Passanger) Update(s string) {
	fmt.Println("Passanger Update", s)
}

func main() {

	bus := Bus{}

	p := &Passanger{}
	s1 := &Station{}
	s2 := &Station{}

	bus.Register(p)
	bus.Notify("S01")

	bus.Register(s1)
	bus.Register(s2)

	bus.Notify("S02")

	bus.Remove(s2)

	bus.Notify("S03")

}
