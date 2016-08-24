package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// //inteiro com sinal
	var inteiro8 int8 //valores com sinal de -128 a 127
	var inteiro16 int16 //valores com sinal de  -32768 a 32767
	// var inteiro32 int32 //valores com sinal de -2147483648 a 2147483647
	// var inteiro int64 //valores com sinal de -9223372036854775808 a 9223372036854775807

	// //inteiro sem sinal
	// var uinteiro8 uint8 //valores sem sinal de 0 a 255
	// var uinteiro16 uint16 //valores sem sinal de 0 a 65535
	// var uinteiro32 uint32 //valores sem sinal de 0 a 4294967295
	// var uinteiro64 uint64 //valores sem sinal de 0 a 18446744073709551615

	// //Alias
	// var inteiroByte byte //alias para uint8
	// var inteiroRune rune //alias para int32

	// //Tipos dependente de plataformas
	// var inteiroUint uint //32 ou 64
	// var inteiroint int //32 ou 64

	// var inteiroUintptr uintptr //inteiro sem sinal suficiente para guardar um valor de ponteiro

	inteiro8 = 5
	inteiro16 = 32762

	fmt.Println(int32(inteiro16) * int32(inteiro16))
	fmt.Println(int16(inteiro8) + inteiro16)

	var teste32 int64 = 10
	var testeLinux int = 999999999999999999

	fmt.Println(int(teste32) + testeLinux)


	fmt.Println(unsafe.Sizeof(testeLinux), unsafe.Sizeof(teste32))
}