package main

import "fmt"

// Interfaces
func main() {

	e := Estructura{}
	fmt.Println(e.MiFuncion())
	fmt.Println(e.Calcular(10, 15))

}

type ejemploInterfaz interface {
	miFuncion() string
	calcular(n1, n2 int) int
}

type Estructura struct {
}

func (*Estructura) MiFuncion() string {
	return "Estoy cimplementado el metodo de la interfaz"
}

func (*Estructura) Calcular(n1, n2 int) int {
	return n1 + n2
}

/*
// estructuras
func main() {

	estructura := Persona{
		Id:     1,
		Nombre: "Sergio",
		Edad:   25,
		Correo: "sfs@gmail.com",
	}
	fmt.Println(estructura)
	fmt.Println("---------------Imprimir estrucura coon formato-----------------")
	fmt.Printf("%+v\n", estructura)
	fmt.Printf("%#v\n", estructura)
	fmt.Println(reflect.TypeOf(estructura))
	fmt.Println("---------------Forma 2 de crear estrucuras-----------------")
	p := new(Persona)
	p.Id = 2
	p.Nombre = "Juan"
	p.Edad = 30
	p.Correo = "juna@gmail.com"
	fmt.Println(p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Println(reflect.TypeOf(p))
	// en genral la primera forma es utuilizada en ORM para crear registros y la segunda forma para podificar registros.

	fmt.Println("---------------estructuras anidadas-----------------")
	categoria := Categoria{Id: 1, Nombre: "Categoria 1", Slug: "categoria-1"}
	producto := Producto{Id: 1, Nombre: "Mesa de computador", Slug: "mesa-de-computador", Precio: 1234, CategoriaId: categoria}
	fmt.Printf("%+v \n", producto)

}

type Persona struct {
	Id     int
	Nombre string
	Edad   int
	Correo string
}

type Categoria struct {
	Id     int
	Nombre string
	Slug   string
}

type Producto struct {
	Id          int
	Nombre      string
	Slug        string
	Precio      int
	CategoriaId Categoria
}

*/

/*
// Recursividad
// es cuando vamos a llamar un procedimineto en si mismo
// ó por ejemplo vamos a llamar un funcion dentro de si misma
func main() {
	miFuncion(0)
}

func miFuncion(valor int) {
	dato := valor + 1
	fmt.Println(dato)
	if dato < 10 {
		miFuncion(dato)
	}
}
*/

/*
// Goroutines
func main() {
	// ejemplo 1
	fmt.Println(miFuncion("Sergio"))
	time.Sleep(time.Second * 5)
	fmt.Println(miFuncion("Paula"))
	// ejemplo 2
	miCalna := make(chan string)
	go func() {
		miCalna <- miFuncion("Pedro")
	}()
	fmt.Println(<-miCalna)
	fmt.Println("Continuear con la ejecución")

}

func miFuncion(nombre string) string {
	return "hola" + nombre
}
*/

/*
// Funciones
func main() {
	miFuncion()
	miFuncionConParametros(5, 10)
	fmt.Println(miFuncionConRetorno("Sergio"))
	nombre, apellido, edad := miFuncionConRetornoMultiple()
	fmt.Printf("Hola %v %v, tu edad es de %v años \n", nombre, apellido, edad)
	fmt.Println("La suma es = ", suma(10, 12))
	Tabla := tabla(2)
	for i := 0; i < 11; i++ {
		fmt.Printf(" 2 x %v = %v \n", i, Tabla())
	}
}

func miFuncion() {
	fmt.Println("Hola mundo")
}

func miFuncionConParametros(n1 int, n2 int) {
	resultado := n1 + n2
	fmt.Println("La suma es: ", resultado)
}

func miFuncionConRetorno(nombre string) string {
	return "tu nombre es: " + nombre
}

func miFuncionConRetornoMultiple() (string, string, int) {
	return "Sergio", "Fernandez", 27
}

// funcion anonima
var suma = func(numero1, numero2 int) int {
	return numero1 + numero2
}

// clousure  -> son funciones que retornan otra funcion
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		resultado := numero * secuencia
		secuencia++
		return resultado
	}
}
*/

/*
// Map
func main() {
	// {"id": 1, "name": "sergio"}
	paises := make(map[string]int)
	paises["Colombia"] = 4000000
	paises["Brazil"] = 2000000
	paises["Argentina"] = 1000000
	paises["Chile"] = 500000
	fmt.Println(paises)
	fmt.Println(reflect.TypeOf(paises))
	//para imprimir la catidad de un pais en especifico
	fmt.Println(paises["Colombia"])
	fmt.Println("--------------------------------")
	paises2 := map[int]string{
		1: "Colombia",
		2: "Brazil",
		3: "Argentina",
		4: "Chile",
	}
	fmt.Println(paises2)
	// imprimir un pais
	fmt.Println(paises2[1])
	fmt.Println("--------------------------------")

	// Validar si existe un valor en el map
	pais, existe := paises2[5]
	if existe {
		fmt.Println(pais)
	} else {
		fmt.Println("No existe")
	}

	// eliminar un elemento de un pais
	fmt.Println("--------------------------------")
	delete(paises2, 1)
	fmt.Println(paises2)

	// recorrer un map con el for
	for id, pais := range paises2 {
		fmt.Println("id: ", id, "nombre: ", pais)
	}

	// ejemplo de un mapa muy utilizado
	// en la creación de apis para manejar las respuesta de las trasacciones  que se dan sobre la Api
	respuesta := map[string]string{
		"estado":  "Ok",
		"mensaje": "mensaje descriptivo",
	}
	fmt.Println(respuesta)

}
*/

/*
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
*/

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
