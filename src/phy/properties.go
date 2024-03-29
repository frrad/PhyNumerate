package phy

import (
//"fmt"
)

func (p *Phy) Score(signal []bool) (score int) {
	key := keygen(signal)
	if answer, ok := p.scorecache[key]; ok {
		return answer
	}

	set(p, signal)

	score = countM(p.root)
	zero(p)

	p.scorecache[key] = score

	return
}

//Returns the score on p relative to probabilities generated from q
func (p *Phy) ScoreRel(q *Phy) *NPoly {

	if p.Size() != q.Size() {
		panic("Can't compute: size mismatch")
	}

	answer := NewNPoly((2 * p.Size()) - 3)

	limit := 1 << uint(p.Size()-1)

	for i := 0; i < limit; i++ {
		test := append(binary(p.Size()-1, i), true)

		score := p.Score(test)

		prob := q.Prob(test)

		prob.Scale(score)
		answer.Add(prob)

	}

	return answer

}

//sets leaves according to signal
func set(p *Phy, signal []bool) {
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
}

//Does majority rule, counts changes
func countM(t *tree) int {
	total := 0
	if t.left.value == 0 {
		total += countM(t.left)
	}
	if t.right.value == 0 {
		total += countM(t.right)
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

//number of adjacent nodes differing in value
func countD(t *tree) int {
	score := 0
	if t.left != nil {
		score += countD(t.left)
		if t.left.value != t.value {
			score++
		}
	}
	if t.right != nil {
		score += countD(t.right)
		if t.right.value != t.value {
			score++
		}
	}
	return score
}

func keygen(signal []bool) uint64 {
	answer := uint64(0)
	for i := 0; i < len(signal); i++ {
		if signal[i] {
			answer += 1 << uint(i)
		}
	}
	return answer
}

func (p *Phy) Prob(signal []bool) *NPoly {
	key := keygen(signal)
	if answer, ok := p.probcache[key]; ok {
		return answer.clone()
	}

	inside := inner(p.root)
	answer := NewNPoly((2 * p.Size()) - 3)
	set(p, signal)

	//list "inside" includes the roots, which we shall omit
	numInside := len(inside) - 1
	combos := 1 << uint(numInside)

	for i := 0; i < combos; i++ {
		test := bipolar(numInside, i)
		//	fmt.Println(test)

		k := 0
		for j := 0; j < numInside; j++ {
			if inside[k] == p.root {
				k++
			}
			inside[k].value = test[j]
			k++
		}

		//peek(p)

		differences := countD(p.root) - 2
		if p.root.left.value != p.root.right.value {
			differences++
		}
		//	fmt.Println(differences)
		answer.Increment(differences, 1)

	}

	zero(p)

	p.probcache[key] = answer

	return answer.clone()
}

//the indexth  (bipolar?) value with nodes digits
func bipolar(nodes, index int) []int {
	if nodes == 1 {
		if index%2 == 0 {
			return []int{-1}
		} else {
			return []int{1}
		}
	}
	if index%2 == 0 {
		return append(bipolar(nodes-1, index/2), -1)
	}

	return append(bipolar(nodes-1, index/2), 1)

}

func binary(height, index int) []bool {
	if height == 1 {
		if index%2 == 0 {
			return []bool{false}
		} else {
			return []bool{true}
		}
	}
	if index%2 == 0 {
		return append(binary(height-1, index/2), false)
	}

	return append(binary(height-1, index/2), true)

}

//returns a slice of pointers to internal decendents of "t"
func inner(t *tree) []*tree {
	if t.left == nil && t.right == nil {
		return nil
	}
	return append(append(inner(t.left), t), inner(t.right)...)

}
