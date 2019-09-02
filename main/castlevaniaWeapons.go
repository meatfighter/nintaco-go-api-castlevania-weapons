package main

import "github.com/meatfighter/nintaco-go-api/nintaco"

type castlevaniaWeapons struct {
	api nintaco.API
}

func newCastlevaniaWeapons() *castlevaniaWeapons {
	return &castlevaniaWeapons{}
}

func (c *castlevaniaWeapons) launch() {

}

func main() {
	newCastlevaniaWeapons().launch()
}
