package main

import (
	"fmt"
)

func main() {
	var f32 float32 = 1.2222224224612 //6 digitos de precisão
	var f64 = 1.2222                  //15 digitos de precisão para float64
	var c64 complex64 = 12.333453     //numeros complexos para float32
	var c128 complex128 = 12.333453   //numeros complexos para float64

	fmt.Println(f32, f64, c64, c128)

	fmt.Println("f32 * 58.565656 = ", f32*58.565656)
	fmt.Println("f64 * 58.565656 = ", f64*58.565656)

	fmt.Println("c64 * 8768464.23", c64*8768464.23)
	fmt.Println("c128 * 8768464.23", c128*8768464.23)

	fmt.Println("[Real] c128 * 8768464.23 = ", real(c128*8768464.23))

}
