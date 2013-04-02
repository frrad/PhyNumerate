package main

import (
	//"fmt"
	"phy"
)

func main() {
	phy1 := phy.NewPhy(4)
	phy2 := phy.NewPhy(4)
	phy1.Assemble([][2]int{[2]int{0, 1}, [2]int{2, 3}, [2]int{0, 2}})
	phy2.Assemble([][2]int{[2]int{0, 2}, [2]int{1, 3}, [2]int{0, 1}})

	phy1.Print()

	options := []bool{true, false}

	answer := phy.NewNPoly(5)

	a := false
	for _, b := range options {
		for _, c := range options {
			for _, d := range options {
				matress := []bool{a, b, c, d}

				score := phy2.Score(matress)

				p := phy1.Prob(matress)

				p.Scale(score)
				answer.Add(p)
			}
		}
	}

	answer.Print()

}
