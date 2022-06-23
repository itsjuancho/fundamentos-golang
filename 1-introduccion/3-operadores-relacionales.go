package main

import "fmt"

func main() {
	/*
		Operadores relacionales:
			Mayor (>)
			Menor (<)
			Mayor o igual (>=)
			Menor o igual (<=)
			Igual (==)
			Diferente (!=)
	*/
	edad := 10
	mayoriaEdad := 18
	var puedeVotar = edad > mayoriaEdad

	nombre := "Juan"
	nombreAmigo := "juan"
	var sonCadenasIguales = nombre == nombreAmigo // No son la misma cadena debido a que deben de ser estrictamente iguales (case-sensitive y espacios)

	fmt.Println("Puede votar?:", puedeVotar)
	fmt.Println("Son la misma cadena?:", sonCadenasIguales)
}
