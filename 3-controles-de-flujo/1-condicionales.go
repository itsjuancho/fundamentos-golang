package main

import "fmt"

func main() {
	var edad int
	fmt.Print("Ingresa tu edad: ")
	fmt.Scanf("%d", &edad)

	//var puedeVotar bool = true
	if edad >= 18 {
		fmt.Println("Puedes votar!")
	} else {
		fmt.Println("No puedes votar!")
	}
}
