package main

import (
	"fmt"

	"github.com/meatfighter/nintaco-go-api/nintaco"
)

const (
	addressLives     = 0x002A
	addressShots     = 0x0064
	addressWhip      = 0x0070
	addressHearts    = 0x0071
	addressSubweapon = 0x015B
)

type castlevaniaWeapons struct {
	api           nintaco.API
	buttonPressed bool
}

func newCastlevaniaWeapons() *castlevaniaWeapons {
	return &castlevaniaWeapons{}
}

func (c *castlevaniaWeapons) launch() {
	c.api.AddFrameListener(c)
	c.api.AddStatusListener(c)
	c.api.AddActivateListener(c)
	c.api.AddDeactivateListener(c)
	c.api.AddStopListener(c)
	c.api.Run()
}

func (c *castlevaniaWeapons) APIEnabled() {
	fmt.Println("API enabled")
}

func (c *castlevaniaWeapons) APIDisabled() {
	fmt.Println("API disabled")
}

func (c *castlevaniaWeapons) Dispose() {
	fmt.Println("API stopped")
}

func (c *castlevaniaWeapons) StatusChanged(message string) {
	fmt.Printf("Status message: %s\n", message)
}

func (c *castlevaniaWeapons) switchWeapons() {
	c.api.WriteCPU(addressSubweapon, c.getNextSubweapon())
}

func (c *castlevaniaWeapons) getNextSubweapon() int {
	switch c.api.ReadCPU(addressSubweapon) {
	case 0x08:
		return 0x09 // dagger -> cross
	case 0x09:
		return 0x0B // cross -> holy water
	case 0x0B:
		return 0x0D // holy water -> axe
	case 0x0D:
		return 0x0F // axe -> stopwatch
	case 0x0F:
		return 0x0A // stopwatch -> rosary
	default:
		return 0x08 // rosary -> dagger
	}
}

func (c *castlevaniaWeapons) FrameRendered() {
	if c.api.ReadGamepad(0, nintaco.GamepadButtonSelect) {
		if !c.buttonPressed {
			c.buttonPressed = true
			c.switchWeapons()
		}
	} else {
		c.buttonPressed = false
	}
}

func main() {
	newCastlevaniaWeapons().launch()
}
