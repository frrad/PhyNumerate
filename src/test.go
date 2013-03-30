package main

import (
	"fmt"
	"phy"
)

func main() {
	phy1 := phy.NewPhy(4)
	phy1.Assemble([][2]int{[2]int{0, 1}, [2]int{2, 3}, [2]int{0, 2}})

	phy1.Print()

	options := []bool{true, false}

	a := true
	for _, b := range options {
		for _, c := range options {
			for _, d := range options {
				matress := []bool{a, b, c, d}
				fmt.Println(phy1.Score(matress), matress)

			}
		}
	}

}
