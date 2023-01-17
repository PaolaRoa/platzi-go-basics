package main

import "fmt"

func main() {
	//go build src/main.go
	//go run src/main.go
	//Declaracion constante
	const pi float64 = 3.14

	const pi2 = 3.14

	fmt.Println(pi, pi2)

	//declaracion variables
	//enteras
	//no declarada anteriormente := crea y asigna
	base := 12
	var altura int = 14
	//Zero values -> asigna valor por defecto
	var area int
	var a float64
	var b string
	var c bool

	fmt.Println(base, altura, area)

	fmt.Println(area, a, b, c)

	//area
	const basecuadrado = 10
	areaCuadrado := basecuadrado * basecuadrado

	fmt.Println(areaCuadrado)

}
