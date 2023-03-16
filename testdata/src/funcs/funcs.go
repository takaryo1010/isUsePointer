package funcs

type Calc struct{ atai1, atai2 int }

func Add(p *Calc) int { 
	return p.atai1 + p.atai2
}

func NewCalc() Calc {
	return Calc{atai1: 10, atai2: 20}
}
