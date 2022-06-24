package main

import "fmt"

func main() {
	/*
		La función make() nos podrá proporcionar un arreglo definiendo su longitud y capacidad máxima
		La capacidad máxima se sobre-escribe luego de que haya llegado a su límite, aumentando la misma en el doble del valor anterior
		Es decir, si la capacidad máxima es de 10 y se sobrepasa ese límite, ya la capacidad máxima será de 20
	*/

	numeros := make([]int, 3, 4)

	// Asignación de valor mediante índice
	numeros[0] = 15

	// Agregamos nuevos elementos al final del arreglo
	numeros = append(numeros, 15)
	numeros = append(numeros, 78)

	fmt.Println(numeros)
	fmt.Println("La longitud de elementos es de", len(numeros), "y su capacidad máxima es de", cap(numeros))
}
