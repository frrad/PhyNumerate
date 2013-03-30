package phy

import "fmt"

type Phy struct {
	leaves []*tree
	root   *tree
}

type tree struct {
	parent *tree
	value  int
	left   *tree
	right  *tree
}

func (p *Phy) Print() {
	fmt.Println("There are", len(p.leaves), "leaves")

	if p.root != nil {
		fmt.Println("Tree is rooted")
	} else {
		fmt.Println("Tree has no root")
	}

	counter := 1
	for i := 0; i < len(p.leaves); i++ {
		place := p.leaves[i]

		for place != nil && place.value == 0 {
			place.value = counter
			counter++
			place = place.parent
		}
	}

	for i := 0; i < len(p.leaves); i++ {
		place := p.leaves[i]

		for place != nil {
			fmt.Print(place.value, "-")
			place = place.parent
		}
		fmt.Print(")\n")

	}

	zero(p)

}

func zero(p *Phy) {
	if p.root != nil {
		zeroTree(p.root)
	} else {
		for i := 0; i < len(p.leaves); i++ {
			place := p.leaves[i]
			for place != nil {
				place.value = 0
				place = place.parent
			}
		}

	}
}

func zeroTree(t *tree) {
	if t.left != nil {
		zeroTree(t.left)
	}
	t.value = 0
	if t.right != nil {
		zeroTree(t.right)
	}
}

func link(A, B int, p *Phy) {
	a := p.leaves[A]
	b := p.leaves[B]
	for a.parent != nil {
		a = a.parent
	}
	for b.parent != nil {
		b = b.parent
	}
	link := tree{left: a, right: b}
	a.parent = &link
	b.parent = &link

}

func NewPhy(n int) *Phy {
	llist := make([]*tree, n)
	for i := 0; i < n; i++ {
		llist[i] = new(tree)
	}

	newPhy := Phy{leaves: llist}

	return &newPhy

}

func (p *Phy) Assemble(instruct [][2]int) {
	if len(instruct)+1 != len(p.leaves) {
		panic("INSTRUCITONS HAVE WRONG LENGTH")
	}
	for _, join := range instruct {
		link(join[0], join[1], p)
	}
	setroot(p)
}

func setroot(p *Phy) {
	root := p.leaves[0]
	for root.parent != nil {
		root = root.parent
	}
	p.root = root
}
