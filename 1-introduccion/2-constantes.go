package main

import "fmt"

/*
	- La constante se asigna fuera de la función principal
	- Se asigna en tiempo de ejecución
	- Se asignan por fuera de la función principal (main)
*/
const CODIGO_CURSO string = "CF589"
const CIUDAD_NATAL string = "Medellín"
const LIMITE_CANTIDAD_INTENTOS int = 3

func main() {
	fmt.Println(CODIGO_CURSO, CIUDAD_NATAL, LIMITE_CANTIDAD_INTENTOS)
}
