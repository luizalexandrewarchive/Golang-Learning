package main

import (
	"fmt"
	"math"
)

func main() {

	var valorMinimo, valorMaximo int

	fmt.Print("Digite o valor minimo e maximo do expoente: ")

	fmt.Scanln(&valorMinimo)
	fmt.Scanln(&valorMaximo)

	exp(valorMinimo, valorMaximo)

}

func exp(init int, end int) {
	if init <= end {
		fmt.Print(math.Pow(float64(2), float64(init)), " ")
		exp(init+1, end)
	}
}
