package main

import "fmt"

// Asignación de constantes múltiples (para evitar poner const repetidas veces)
const (
	DOMINGO   int = 0
	LUNES     int = 1
	MARTES    int = 2
	MIERCOLES int = 3
	JUEVES    int = 4
	VIERNES   int = 5
	SABADO    int = 6
)

/*
	También existe una manera más sencilla aún. Y es con el valor iota.
	¿Qué es el valor iota? Es una palabra reservada que permite especificar una secuencia de valores automática.
	Por defecto, se encontrará en 0. Así el lenguaje ya sabe, que las constantes siguientes se incrementarán posteriormente en +1
*/

const (
	CERO int = iota // Si se quiere modificar su valor de inicio, pondremos "iota + valor de inicio"
	PRIMERO
	SEGUNDO
	TERCERO
	CUARTO
	QUINTO
	SEXTO
)

func main() {
	fmt.Println(CERO)
	fmt.Println(PRIMERO)
	fmt.Println(SEGUNDO)
	fmt.Println(TERCERO)
	fmt.Println(CUARTO)
	fmt.Println(QUINTO)
	fmt.Println(SEXTO)
}
