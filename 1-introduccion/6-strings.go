package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nombreMateria string = "Programación Orientada a Objetos"
	var profesorMateria = "Nico"
	bimestre := "Tercero"

	longitudNombreProfesor := len(profesorMateria) // Retorna la cantidad de caracteres que tiene la cadena. Los espacios en blanco cuentan.

	fmt.Println(longitudNombreProfesor)

	tercerCaracter := bimestre[3]                // Retornará el código ASCII del caracter. Para convertirlo, usaremos reflect.TypeOf()
	miCaracter := reflect.TypeOf(tercerCaracter) // Nos retornará la conversión del caracter, es decir, el tipo de dato

	ultimoCaracter := nombreMateria[len(nombreMateria)-1]

	fmt.Println("El tercer caracter de la variable bimestre es:", miCaracter)
	fmt.Printf("%c\n", tercerCaracter) // Para poder visualizar el caracter de manera natural, tendremos que usar Printf en vez de Println
	fmt.Printf("%c", ultimoCaracter)
}
