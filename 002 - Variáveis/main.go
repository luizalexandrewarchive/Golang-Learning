package main

import "fmt"

func main() {

	// bool,string,int,int8,int16,int32,int64
	// uint,uint8,uint16,uint32,uint64,uintptr
	// byte,rune,float32,float64,complex64,complex128
	
	var numero int 
	numero = 25
	fmt.Println(numero + 15)

	nome := "luizalexandrew"
	numero = 12

	var numerx = 122

	fmt.Println(numerx)

	fmt.Println(nome, numero)

	nome, numero = "Google", 1

	fmt.Println(nome, numero)

	nome2, numero2 := "Microsoft", 2

	fmt.Println(nome2, numero2)

	var final string = "BIIIRRR"
	fmt.Println(final)

	fmt.Println(nome, numero, "BIIIRRR")

	var(
		personagen = "Goku"
		inimigo1, inimigo2 = "Cell", "Majin BOO"
		vida = 23
	)	

	fmt.Println(personagen, inimigo1, inimigo2, "vida:", vida)

	fmt.Println("Os inimigos existentes são: " + inimigo1 + inimigo2)

	var flutuante32 float32 = 23.22222222222222222222222222222
	var flutuante64 float64 = 23.22222222222222222222222222222

	fmt.Println(flutuante32, flutuante64)

}

