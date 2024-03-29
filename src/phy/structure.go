package phy

import "fmt"

type Phy struct {
	leaves     []*tree
	root       *tree
	probcache  map[uint64]*NPoly
	scorecache map[uint64]int
}

type tree struct {
	parent *tree
	value  int
	left   *tree
	right  *tree
}

func (p *Phy) Print() {
	fmt.Println("There are", p.Size(), "leaves")

	if p.root != nil {
		fmt.Println("Tree is rooted")
	} else {
		fmt.Println("Tree has no root")
	}

	counter := 1
	for i := 0; i < p.Size(); i++ {
		place := p.leaves[i]

		for place != nil && place.value == 0 {
			place.value = counter
			counter++
			place = place.parent
		}
	}

	for i := 0; i < p.Size(); i++ {
		place := p.leaves[i]

		for place != nil {
			fmt.Print(place.value, "-")
			place = place.parent
		}
		fmt.Print(")\n")

	}

	zero(p)

}

//makes sure we can reach root from each leaf
func (p *Phy) Validate() bool {
	for i := 0; i < len(p.leaves); i++ {
		current := p.leaves[i]
		for current.parent != nil {
			current = current.parent
		}
		if current != p.root {
			return false
		}
	}
	return true
}

//shows contents of tree for debugging purposes
func peek(p *Phy) {

	for i := 0; i < p.Size(); i++ {
		place := p.leaves[i]

		for place != nil {
			fmt.Print(place.value, "<>")
			place = place.parent
		}
		fmt.Print(")\n")

	}

}

func (p *Phy) Size() int {
	return len(p.leaves)
}

func zero(p *Phy) {
	if p.root != nil {
		zeroTree(p.root)
	} else {
		for i := 0; i < p.Size(); i++ {
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
	if a == b {
		return
		panic("Already linked")
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

	newPhy := Phy{leaves: llist, probcache: make(map[uint64]*NPoly), scorecache: make(map[uint64]int)}

	return &newPhy

}

func (p *Phy) Assemble(instruct [][2]int) {
	if len(instruct)+1 != p.Size() {
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
