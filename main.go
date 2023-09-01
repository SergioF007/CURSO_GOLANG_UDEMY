package main

import (
	"fmt"
	"reflect"
)

// reflect y TypeOf
// import "reflect"
func main() {
	string1 := 3424
	fmt.Println(reflect.TypeOf(string1))

}

/*
// tipos de datos
func main() {
	var string1 string = "texto"
	fmt.Println(string1)

	textoGrande := `Aprender un poco cada día marca la diferencia. Hay estudios que muestran que los estudiantes que hacen del aprendizaje un hábito tienen una mayor probabilidad de alcanzar sus objetivos. Reserva tiempo para aprender y recibe recordatorios con la herramienta de planificación del aprendizaje.`
	fmt.Println("------------------------------------------------")
	fmt.Println(textoGrande)
	var estado bool = false
	fmt.Println("------------------------------------------------")
	fmt.Println(estado)
	var flotan64 float64 = 32.33
	fmt.Println("------------------------------------------------")
	fmt.Println(flotan64)
	var flotan32 float32 = 32.32
	fmt.Println("------------------------------------------------")
	fmt.Print(flotan32)
	var entero_int8 int8 = 123 // -128 a 127
	fmt.Println("------------------------------------------------")
	fmt.Println(entero_int8)
	var entero_int16 int16 = 123 // -32768 a 32767
	fmt.Println("------------------------------------------------")
	fmt.Println(entero_int16)
	var entero_int32 int32 = 123 // -2147483648 a 2147483647
	fmt.Println("------------------------------------------------")
	fmt.Println(entero_int32)
	var entero_int64 int64 = 123 // -9223372036854775808 a 9223372036854775807
	fmt.Println("------------------------------------------------")
	fmt.Println(entero_int64)
	var entero_uint8 uint8 = 233 // 0 a 255
	fmt.Println("------------------------------------------------")
	fmt.Println(entero_uint8)
	var entero_uint16 uint16 = 233 // 0 a 65535
	fmt.Println("------------------------------------------------")
	fmt.Println(entero_uint16)
	var entero_uint32 uint32 = 233 // 0 a 4294967295
	fmt.Println("------------------------------------------------")
	fmt.Println(entero_uint32)

}
*/

/*
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
*/
