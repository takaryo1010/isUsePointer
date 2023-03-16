package a

type Calc struct{ atai1, atai2 int }

func (p *Calc) Add() int { // want "use pointer"
	return p.atai1 + p.atai2
}
func (p *Calc) Sub() int { // want "use pointer"
	return p.atai1 - p.atai2
}
func NewCalc() Calc {
	return Calc{atai1: 10, atai2: 20}
}
