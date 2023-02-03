package main

//Paquete para interacturar en consola
import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	//go build src/main.go.go
	//go run src/main.go.go
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

	operations()
	dataTypes()
	fmtExploration()
	normalFunction("hola mundo")
	tripleArguments(1, 2, "hola")
	fmt.Println(returnValue(2))
	value1, value2 := doubleReturn(2)
	fmt.Println(value1, value2)
	//guión al piso ignora valor retornado
	_, value3 := doubleReturn(4)
	fmt.Println(value3)

	cycles()
	conditinals()
	deferBackContinue()
	arraySlices()

}

func operations() {
	x := 10
	y := 50

	//suma
	result := x + y
	fmt.Println("suma", result)

	//resta
	result = x - y
	fmt.Println("resta", result)

	//multiplicación
	result = x * y
	fmt.Println("multi", result)

	//división
	result = x / y
	fmt.Println("división", result)

	//modulo
	result = x % y
	fmt.Println("modulo", result)

	//incremental
	x++
	fmt.Println(x)

	//dremento
	x--
	fmt.Println(x)

}

func dataTypes() {
	//GO identifica el tipo de dato, sin embargo especificarlo
	//a la hora de asginar mejora el rendimiento

	//Enteros int -> toma tipo del SO
	var i int = 8
	//int8, int16, int32, int64
	//Enteros positivos, siempre positivos
	//uint .. 8, 16, 32, 64
	var positivo uint = 12

	//decimales
	var a float32 = 3.1
	var b float64 = 3.2

	//texto
	var s string = "string"

	//booleanos

	var boolean bool = true

	//números complejos
	//complex 64 o 32
	complex := 10 + 8i

	fmt.Println(i, positivo, a, b, s, boolean, complex)

}

func fmtExploration() {
	helloMessage := "hello"
	worldMessage := "world"

	//imprimir linea
	fmt.Println(helloMessage)
	fmt.Println(worldMessage)

	//printF
	//%s string, %d entero, %v tipo desconocido
	nombre := "platzi"
	cursos := 500

	fmt.Printf("%s tiene más de %d cursos \n", nombre, cursos)

	//Sprintf no imprime solo guarda

	message := fmt.Sprintf("%v tiene más de %v cursos", nombre, cursos)
	fmt.Println(message)

	//tipo de dato
	fmt.Printf("%T\n", helloMessage)
}

// funciones
func normalFunction(message string) {
	fmt.Println(message)
}

// func tripleArguments(a int, b int, c string) {
func tripleArguments(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func cycles() {
	//for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//for while
	counter := 10
	for counter < 20 {
		fmt.Println(counter)
		counter++
	}
	// for forever
	//counterForever := 0
	/*for {
		fmt.Println(counterForever)
		counterForever++
	}*/

}

func conditinals() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("es 1")
	} else {
		fmt.Println("no es 1")
	}

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("true")
	}

	//texto a int
	value, err := strconv.Atoi("14")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
	//switch
	modulo := 5 % 2

	switch modulo {
	case 0:
		fmt.Println("es par")
	default:
		fmt.Println("es impar")
	}

	value2 := 20
	switch {
	case value2 > 100:
		fmt.Println("mayor a 100")
	case value2 < 100:
		fmt.Println("menor que 100")
	default:
		fmt.Println("no")

	}

}

func deferBackContinue() {
	//Defer
	//ejecuta ultima funcion antes de que termine la ejecución
	defer fmt.Println("termino")
	fmt.Println("codigo")

	//Continue y break -> usados dentro de for
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//continue sigue con el ciclo
		if i == 2 {
			fmt.Println("es 2")
			continue
		}

		//break rompe la ejecución
		if i == 8 {
			fmt.Println("Break")
			break

		}

	}
}

func arraySlices() {
	//arrays inmutables
	var array [4]int
	array[0] = 3
	array[3] = 8

	// len longitud, cap capacidad
	fmt.Println(array, len(array), cap(array))

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice, len(slice), cap(slice))

	//Métodos slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//añadir elementos
	slice = append(slice, 9)

	fmt.Println(slice)

	slice2 := []int{8, 99, 85}
	slice = append(slice, slice2...)

	fmt.Println(slice)

}
