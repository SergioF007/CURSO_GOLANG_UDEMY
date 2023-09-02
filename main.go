package main

import "fmt"

// Arreglos y Slice
func main() {
	// arreglo (array)
	var paises [4]string
	paises[0] = "Colombia"
	paises[1] = "Peru"
	paises[2] = "Argentina"
	paises[3] = "Uruguay"

	fmt.Println(paises)
	fmt.Println("--------------------------------")

	// slice
	var paises2 = make([]string, 5)
	paises2[0] = "Colombia"
	paises2[1] = "Peru"
	paises2[2] = "Argentina"
	paises2[3] = "Uruguay"
	paises2[4] = "España"
	fmt.Println(paises2)
	fmt.Println("--------------------------------")

	// agregar un valor al slice
	paises2 = append(paises2, "Chile")
	fmt.Println(paises2)
	fmt.Println("El largo del arreglo es: ", len(paises2))

	// eliminar registros de un slice
	paises2 = append(paises2[:5], paises2[5+1:]...)
	fmt.Println("--------------------------------")
	fmt.Println(paises2)

}

/*
// cilos e itereaciones
func main() {
	i := 1
	for i < 11 {
		fmt.Println(i)
		i++
	}
	fmt.Println("--------------------------------")
	for i2 := 1; i2 < 11; i2++ {
		fmt.Println(i2)
	}

}
*/

/*
// Punteros
func main() {
	color := "rojo"
	fmt.Println(color, &color)
}
*/

/*
// reflect y TypeOf
// import "reflect"
func main() {
	string1 := 3424
	fmt.Println(reflect.TypeOf(string1))

}
*/

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
