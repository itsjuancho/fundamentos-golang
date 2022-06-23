package main

import "fmt"

const NACIONALIDAD_PERMITIDA string = "Colombiano"
const MAYORIA_EDAD int = 18

func main() {
	/*
		Operadores lógicos:
			AND (&&)
			OR (||)
			Negación (!)
	*/
	nacionalidadCiudadano := "Colombiano"
	edadCiudadano := 18
	puedeVotar := nacionalidadCiudadano == NACIONALIDAD_PERMITIDA && edadCiudadano >= MAYORIA_EDAD
	fmt.Println("Cumple con las condiciones para votar en Colombia?:", puedeVotar)

	nacionalidadExtranjero := "Panameño"
	edadExtranjero := 17
	puedeVotarPorEdad := nacionalidadExtranjero == NACIONALIDAD_PERMITIDA || edadExtranjero >= MAYORIA_EDAD
	fmt.Println("Cumple con alguna de las condiciones para votar, sin importar el país?:", puedeVotarPorEdad)

	esExtranjero := !(nacionalidadCiudadano == NACIONALIDAD_PERMITIDA)
	fmt.Println("Es extranjero?:", esExtranjero)
}
