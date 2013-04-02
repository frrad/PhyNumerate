package phy

import "fmt"

//polynomial of the form (1-p)^degree + (1-p)^(degree-1)p^1 + ... 
type NPoly struct {
	degree       int //Number of edges
	coefficients []int
}

func NewNPoly(deg int) *NPoly {
	return &NPoly{degree: deg, coefficients: make([]int, deg+1)}
}

func (p *NPoly) Increment(changes, incr int) {
	p.coefficients[changes] += incr
}

//multiplies the polynomial p by scalar lambda
func (p *NPoly) Scale(lambda int) {
	for i := range p.coefficients {
		p.coefficients[i] *= lambda
	}
}

//adds q to p
func (p *NPoly) Add(q *NPoly) {
	if p.degree != q.degree {
		panic("CAN'T ADD: DEGREES DON'T MATCH")
	}
	for i := range p.coefficients {
		p.coefficients[i] += q.coefficients[i]
	}
}

func (p *NPoly) Print() {
	for i := 0; i <= p.degree; i++ {
		a := p.coefficients[i]
		b := p.degree - i
		c := i
		if a != 0 {
			if a != 1 {
				fmt.Print(a, "*")
			}
			if b != 0 {
				if b == 1 {
					fmt.Print("(1-p)")
				} else {
					fmt.Print("(1-p)^", p.degree-i)
				}
			}

			if b >= 1 && c >= 1 {
				fmt.Print("*")
			}

			if c != 0 {
				if c == 1 {
					fmt.Print("p")
				} else {
					fmt.Print("p^", c)
				}
			}

			fmt.Print("+")

		}
	}
	fmt.Print("\b")
	fmt.Println(" ")
}
