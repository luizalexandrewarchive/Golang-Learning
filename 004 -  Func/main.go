package main

import (
	"fmt"
	"os"
	"strconv"
)

func validarArgs() {

	if len(os.Args) < 3 {
		fmt.Println("Passe o <valores> e <unidade>")
		os.Exit(0)
	}

}

func getUnit() string{

	return os.Args[len(os.Args)-1]

}

func getValues() []float32{

	return toFloat(os.Args[1 : len(os.Args)-1])

}

func toFloat(valoresOrigem []string) []float32{	

	var valoresFloat []float32

	for i, v := range valoresOrigem {

		valor, err := strconv.ParseFloat(v,64)		

		if err == nil {
			valoresFloat = append(valoresFloat, float32(valor))			
		}else{
			fmt.Printf("O valor %s na posição %d não pode ser convertido!\n", v, i)
		}

	}
	return valoresFloat

}

func convert(unidadeOrigem string, valoresOrigem []float32) {
	
	if unidadeOrigem == "CtoF" {
		fmt.Println("Celsius to Fahrenheit")
		celsiusToFahrenheit(valoresOrigem)
	} else if unidadeOrigem == "FtoC" {
		fmt.Println("Fahrenheit to Celsius")
		fahrenheitToCelsius(valoresOrigem)
	}else{
		fmt.Println(unidadeOrigem, "não é uma unidade válida.")
		os.Exit(0)
	}	
	
}

func fahrenheitToCelsius(valoresOrigem []float32) {

	for i, v := range valoresOrigem {
		fmt.Println(i, " - ", v, "°F", " == " , (v-32)/1.8 , "C°")
	}

}


func celsiusToFahrenheit(valoresOrigem []float32) {	

	for i, v := range valoresOrigem {
		fmt.Println(i, " - ", v, "°C", " == " , v*1.8 + 32 , "F°")
	}
	
}



func main(){
	
	validarArgs()

	var unidadeOrigem string = getUnit()
	var valoresOrigem []float32 = getValues()

	convert(unidadeOrigem, valoresOrigem)

}


