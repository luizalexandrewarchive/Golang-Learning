package main

import (
	"fmt"
)

type Pato struct {
	nome string
}

func (p Pato) saudacao() {
	fmt.Println("Oi, eu sou o pato", p.nome, "e faço")
}

func NewPato() Pato {
	d := new(Pato)
	d.nome = "Pato"
	return *d
}

func main() {
	pato := NewPato()
	pato.saudacao()
}
