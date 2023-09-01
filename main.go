package main

import "fmt"

// variables y constantes
const MiConstante = "MiConstante" // public

func main() {

	// declaracion por inferencia
	var nombre string
	fmt.Println(nombre)
	// declaracion rapida o corta
	nombre2 := "sergio"
	fmt.Println(nombre2)
	fmt.Printf("El valor de mis constante es: %s\n", nombre)
}
