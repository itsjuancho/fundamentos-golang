package main

import "fmt"

func main() {
	var numeros [5]int // El valor y/o cantidad de elementos que puede almacenar, no podrá ser modificado ni ampliado posteriormente
	fmt.Println(numeros)

	numeros[0] = 15
	numeros[1] = 30
	numeros[2] = 45
	numeros[3] = 60
	numeros[4] = 75
	// Si asignamos un valor a un índice inexistente o asignar un tipo de dato diferente el declarado, obtendremos un error

	fmt.Println(numeros)

	// Si queremos inicializarlo de una manera sencilla, se hace de la siguiente manera
	edades := [6]int{5, 89, 69, 56, 85, 32}
	fmt.Println(edades)

	/*
		Si deseamos que se inicialice un arreglo sin una cantidad especifica de elementos
		podremos hacer uso del operador "..." para indicar que el arreglo es dinámico
		y que su espacio dependerá de la cantidad de elementos
	*/
	personas := [...]string{"Juan", "Sofía", "Mariana", "Sebastián", "Lesly"}
	fmt.Println(personas)
	fmt.Println("La longitud del arreglo personas es de:", len(personas))

	/*
		Para alterar el orden de cada elemento de un array, podemos modificar su indice.
		Sin embargo, si no asignamos un indice que sea menor a la longitud del arreglo,
		el lenguaje le asignará automáticamente un espacio vacío. Inclusive, si asignamos a un elemento
		un indice mayor a la longitud, los indices que no posean valor tendrán el mismo efecto,
		es decir, se les asignará un valor vacío
	*/
	monedas := [...]string{0: "COP", 1: "MXN", 2: "USD", 3: "ARG", 4: "EUR"}
	fmt.Println(monedas)

	subMonedas := monedas[0:3] // De esta manera, estamos generando un subarreglo (en realidad, hacemos un slice)
	fmt.Println(subMonedas)    // En realidad imprime desde 0 hasta el límite indicado - 1
}
