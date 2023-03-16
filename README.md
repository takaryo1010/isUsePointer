# isUsePointer

`isUsePointer` checks if the receiver of a method in your code is a pointer receiver mixed with a value receiver.
```go
package methods

type Calc struct{ atai1, atai2 int }

func (p *Calc) Add() int { // want "use pointer"
	return p.atai1 + p.atai2
}
func (p *Calc) Sub() int { // want "use pointer"
	return p.atai1 - p.atai2
}
func (p Calc) Mul() int { // want "Mixed use and non-use of pointers"
	return p.atai1 * p.atai2
}

func NewCalc() Calc {
	return Calc{atai1: 10, atai2: 20}
}
```
If the receiver is a pointer receiver, it outputs "use pointer".
If the receiver is a value receiver, it outputs "not use pointer".
If the receivers are mixed up, it outputs "Mixed use and non-use of pointers" at that point, and every time there is a receiver thereafter, it outputs "Mixed use and non-use of pointers".
## Install

```sh
$ go get github.com/takaryo1010/isUsePointer/cmd/isUsePointer
```

## Usage

```sh
$ go vet -vettool=`which isUsePointer` pkgname
```
