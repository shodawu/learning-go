package hubs

import "number-game/interfaces"

// Controller ...
type Controller struct {
	Hubs []interfaces.ICommand
}

// AddCommand ...
func (c *Controller) AddCommand(iCmd interfaces.ICommand) {
	c.Hubs = append(c.Hubs, iCmd)
}
