package main

import "fmt"

func main() {
	var nombre string
	var edad int
	var altura float32

	/*
		Scanf es el método quizá, más limitado que existe para leer valores de entrada.
		Esto, se debe a que solo lee un String sin espacios en blanco.
		Es por ello, que hay más alternativas y más eficientes, sin tantas limitaciones
	*/

	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s\n", &nombre)

	fmt.Print("Ingresa tu edad en números: ")
	fmt.Scanf("%d\n", &edad)

	fmt.Print("Ingresa tu altura con punto decimal: ")
	fmt.Scanf("%f", &altura)

	fmt.Printf("Hola %s con edad %d y una altura %.2f", nombre, edad, altura)

}
