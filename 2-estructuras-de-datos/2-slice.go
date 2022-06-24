package main

import "fmt"

func main() {

	/*
		Diferencias entre los Slices y Arreglos:
			- Son dinámicos
			- Pueden cambiar su longitud en tiempo de ejecución (++ o --)

		- Están estrechamente relacionados con los arreglos. Inclusive, los slices no existen sin los arreglos (básicamente)
	*/
	numeros := []int{10, 20, 30, 40, 50, 60} // Referencia a un arreglo base. Cualquier cambio que se haga en el arreglo base, se verán reflejados en los slice
	numeros = append(numeros, 70)            // Append es para añadir al final un nuevo elemento al arreglo. Este, retorna un nuevo arreglo, se puede asignar al mismo arreglo
	numeros = append(numeros, 80)
	numeros = append(numeros, 90)

	nuevoSlice := numeros[0:5]
	segundoSlice := numeros[0:3]

	numeros[0] = 200 // Cambia en el arreglo base y automáticamente se ve reflejado el cambio en los slices

	fmt.Println(numeros)
	fmt.Println(nuevoSlice)
	fmt.Println(segundoSlice)

	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre"}

	longitudMeses := len(meses) // Cantidad de elementos almacenados en el arreglo
	capacidad := cap(meses)     // Cantidad máxima de elementos que nuestro slice puede almacenar

	fmt.Printf("La longitud del arreglo meses es %v, y la capacidad es de %v - %p\n", longitudMeses, capacidad, meses)
	// %p => obtener la referencia en memoria

	meses = append(meses, "Octubre") // Si la estructura se encuentra en su capacidad máxima, se generará un nuevo arreglo
	meses = append(meses, "Noviembre")
	meses = append(meses, "Diciembre")
	longitudMeses = len(meses)
	capacidad = cap(meses)
	fmt.Printf("La longitud del arreglo meses es %v, y la capacidad es de %v - %p", longitudMeses, capacidad, meses)

	/* ==================== RESUMEN ==================== */
	grados := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	inicio := grados[0:3]
	final := grados[3:6]
	fmt.Println(inicio, final)
}
