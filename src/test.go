package main

import (
	"fmt"
	"math/rand"
	"phy"
)

func main() {
	testlength := 10
	testquantity := 25
	testes := make([]*phy.Phy, testquantity)

	for i := 0; i < testquantity; i++ {
		testes[i] = phy.NewPhy(testlength)

		for !testes[i].Validate() {
			testes[i] = phy.NewPhy(testlength)

			buildcode := make([][2]int, 0)
			for j := 0; j < testlength-1; j++ {
				buildcode = append(buildcode, [2]int{rand.Int() % testlength, rand.Int() % testlength})
			}
			testes[i].Assemble(buildcode)
		}

		fmt.Println("\nNumber", i)
		testes[i].Print()
	}

	for i := 0; i < testquantity; i++ {

		for j := 0; j < testquantity; j++ {

			fmt.Println("Scored on ", i, "with probs from", j)
			testes[i].ScoreRel(testes[j]).Print()
		}
	}
}
