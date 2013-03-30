package phy

func (p *Phy) Score(signal []bool) (score int) {
	if len(signal) != p.Size() {
		panic("Size mismatch")
	}
	if p.root == nil {
		panic("Tree incomplete")
	}
	for i := 0; i < p.Size(); i++ {
		if signal[i] {
			p.leaves[i].value = 1
		} else {
			p.leaves[i].value = -1
		}
	}

	score = count(p.root)
	zero(p)
	return
}

func count(t *tree) int {
	total := 0
	if t.left.value == 0 {
		total += count(t.left)
	}
	if t.right.value == 0 {
		total += count(t.right)
	}

	//Agreement
	if t.left.value == t.right.value {
		t.value = t.left.value
		return total
	}

	//Disagreement
	if t.left.value+t.right.value == 0 {
		t.value = 2
		total++
		return total
	}

	//One neutral, one not
	t.value = t.left.value + t.right.value - 2
	return total

}
