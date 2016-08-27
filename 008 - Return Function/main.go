package main

import (
	"fmt"
	"math"
)

func main() {
	var quadro = potencia(2)
	var dez = potencia(10)
	fmt.Println(quadro(25))
	fmt.Println(dez(25))
	fmt.Println(dez(1))
	fmt.Println(dez(3))
}

func potencia(exp int) func(int) int {
	return func(base int) int {
		return int(math.Pow(float64(base), float64(exp)))
	}
}
