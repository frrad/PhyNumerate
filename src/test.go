package main

import (
	"phy"
)

func main() {
	phy1 := phy.NewPhy(4)
	phy1.Assemble([][2]int{[2]int{0, 1}, [2]int{2, 3}, [2]int{0, 2}})

	phy1.Print()

}
