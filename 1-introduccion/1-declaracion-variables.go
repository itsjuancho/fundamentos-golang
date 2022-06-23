package main

import "fmt"

/*
	go build archivo.go = genera un archivo ejecutable
	go run archivo.go = ejecuta el código sin crear un archivo ejecutable
*/

func main() {
	fmt.Println("Hola Mundo, desde Golang!")

	// En las variables no será posible re-asignarles un valor que no corresponda a su tipo declarado
	var nombre string // Por defecto inician en vacío = ""
	var edad int      // Por defecto un integer inicia en 0

	nombre = "Juan Andrés"
	edad = 19

	// Se puede realizar una declaración implícita (sin indicar el tipo de dato). Y se puede reasignar normalmente
	apellidos := "Pérez"

	// También puede declararse así
	var altura = 1.76

	// La función Println recibe N cantidad de parámetros, así que pueden colocarse todas las variables juntas sin hacer múltiples repeticiones de la función
	fmt.Println(nombre, apellidos, edad, altura)
	fmt.Println("Tu nombre es:", nombre, "y tus apellidos son", apellidos, "con una edad de", edad, "y estatura de", altura)

	// Declaración de múltiples variables
	var ciudad, pais string = "Medellín", "Colombia"
	fmt.Println(pais, ciudad)

	// Declaración de múltiples variables de un distinto tipo
	numeroTelefono, cantidadTelefonos := "85964789", 2
	fmt.Println(numeroTelefono, cantidadTelefonos)
}
