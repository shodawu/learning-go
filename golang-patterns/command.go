package main

import "fmt"

// 電視機的面板有「Power」、「Vol+」、「Vol-」、「Ch+」、「Ch-」等按鈕
// 遙控器上有10個按鈕
// 要使用遙控器控制電視機

type Power struct {
	ICommand
	IsOn bool
}

func (p *Power) Exec() {
	p.PressPower()
}
func (p *Power) PressPower() {
	p.IsOn = !p.IsOn
	fmt.Println("Power now is ", p.IsOn)
}

type TVConsole struct {
	Pwr Power
}

type ICommand interface {
	Exec()
}

type DefCommand struct {
	ICommand
}

func (df *DefCommand) Exec() {
	fmt.Println("Do nothing!!")
}

type Controller struct {
	Buttons []ICommand
}

func (c *Controller) ResetButtons() {
	c.Buttons = []ICommand{}
	i := 0
	for i < 10 {
		c.Buttons = append(c.Buttons, &DefCommand{})
		i++
	}
}

func (c *Controller) SetupButton(iBtn int, iCmd ICommand) {
	c.Buttons[iBtn] = iCmd
}

func main() {
	tv := TVConsole{}
	tv.Pwr.PressPower()
	tv.Pwr.PressPower()

	c := Controller{}
	c.ResetButtons()

	fmt.Println(len(c.Buttons))
	c.Buttons[0].Exec()

	c.SetupButton(0, &tv.Pwr)
	c.Buttons[0].Exec()
	c.Buttons[0].Exec()
}
