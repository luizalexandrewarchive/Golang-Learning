package main

import (
	"fmt"
	"math"
)

func main() {
	var potenciaDez = potencia(10)
	var potenciaDois = potencia(2)
	fmt.Println(potenciaDez(10))
	fmt.Println(potenciaDois(5))
}

func potencia(exp int) func(int) int {
	return func(base int) int {
		return int(math.Pow(float64(base), float64(exp)))
	}
}
