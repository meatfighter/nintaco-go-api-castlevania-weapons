package main

import (
	"fmt"

	"github.com/meatfighter/nintaco-go-api/nintaco"
)

const (
	addressLives1     = 0x002A
	addressLives2     = 0x2097
	addressLives3     = 0x2098
	addressHitPoints1 = 0x0044
	addressHitPoints2 = 0x0045
	addressHitPoints3 = 0x2067
	addressShots1     = 0x0064
	addressShots2     = 0x0141
	addressWhip       = 0x0070
	addressHearts1    = 0x0071
	addressHearts2    = 0x2077
	addressHearts3    = 0x2078
	addressSubweapon  = 0x015B
)

type castlevaniaWeapons struct {
	api           nintaco.API
	buttonPressed bool
}

func newCastlevaniaWeapons() *castlevaniaWeapons {
	return &castlevaniaWeapons{
		api: nintaco.GetAPI(),
	}
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

func (c *castlevaniaWeapons) swapWeapons() {

	// get max hit points
	c.api.WriteCPU(addressHitPoints1, 0x40)
	c.api.WriteCPU(addressHitPoints2, 0x40)
	for i := 7; i >= 0; i-- {
		c.api.WritePPU(addressHitPoints3+i, 0xDA)
	}

	// get 99 lives
	c.api.WriteCPU(addressLives1, 99)
	c.api.WritePPU(addressLives2, 0xD9)
	c.api.WritePPU(addressLives3, 0xD9)

	// get triple shot
	c.api.WriteCPU(addressShots1, 2)
	c.api.WriteCPU(addressShots2, 1)

	// get long chain whip
	c.api.WriteCPU(addressWhip, 2)

	// get 99 hearts
	c.api.WriteCPU(addressHearts1, 99)
	c.api.WritePPU(addressHearts2, 0xD9)
	c.api.WritePPU(addressHearts3, 0xD9)

	// swap subweapons
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
		return 0x08 // none/rosary -> dagger
	}
}

func (c *castlevaniaWeapons) FrameRendered() {
	if c.api.ReadGamepad(0, nintaco.GamepadButtonSelect) {
		if !c.buttonPressed {
			c.buttonPressed = true
			c.swapWeapons()
		}
	} else {
		c.buttonPressed = false
	}
}

func main() {
	nintaco.InitRemoteAPI("localhost", 9999)
	newCastlevaniaWeapons().launch()
}
