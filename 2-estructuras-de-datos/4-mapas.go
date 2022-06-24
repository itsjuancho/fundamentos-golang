package main

import "fmt"

func main() {
	/*
		Los mapas son un tipo de estructura de datos desorganizada que consiste en un tipo de dato llave-valor
		Hace las veces de lo que comunmente conocemos como arrays asociativos
		Para ello, tenemos la siguiente estructura
			make(		=> la función para crear el mapa
				map		=> palabra reservada para especificar que nuestra estructura a crear será un map
				[int]	=> el tipo de dato que será la llave o índice del elemento
				string	=> el tipo de dato correspondiente al valor de la llave
			)
	*/
	dias := make(map[int]string)

	dias[0] = "Domingo"
	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[3] = "Miércoles"
	dias[4] = "Jueves"

	// Para eliminar una llave (y por consiguiente, su respectivo valor) usaremos la función delete, pasándole como argumentos
	// la variable que contiene el map y el valor de la llave a eliminar
	delete(dias, 1)

	fmt.Println(dias)

	estudiantes := make(map[string][]int)
	estudiantes["Andres"] = []int{5, 3, 4, 2}
	nombreNuevaEstudiante := "Laura"
	estudiantes[nombreNuevaEstudiante] = []int{5, 3, 4, 2}
	fmt.Println(estudiantes)

	/* ========================== ITERAR SOBRE MAPAS ========================== */

	usuarios := map[int]string{} // Incializado en vacío

	usuarios[1] = "Usuario 1"
	usuarios[2] = "Usuario 2"
	usuarios[3] = "Usuario 3"

	for llave, valor := range usuarios {
		fmt.Println("Key:", llave, "| Value:", valor)
	}
}
