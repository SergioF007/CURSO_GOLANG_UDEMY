package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("La hora actual es: ", time.Now())
	fecha := time.Now()
	fmt.Println("El año es: ", fecha.Year())
	fmt.Println("El mes es: ", int(fecha.Month()))
	fmt.Println("el numero del mes: ", int(fecha.Month()))
	fmt.Println("El dia es: ", fecha.Day())
	fmt.Println("La hora es: ", fecha.Hour())
	fmt.Println("Los minutos son: ", fecha.Minute())
	fmt.Println("Los segundos son: ", fecha.Second())
	fmt.Printf("El tipo es: %T \n", fecha)
	fmt.Printf("%v %v de %v de %v a las %v:%v:%v \n", fecha.Weekday(), fecha.Day(), fecha.Month(), fecha.Year(), fecha.Hour(), fecha.Minute(), fecha.Minute())
	fmt.Println("--------------Operaciones con fecha------------------")
	ahora := time.Now()
	fmt.Println("Fecha en este momento")
	fmt.Printf("%v/%v/%v \n", ahora.Day(), ahora.Month(), ahora.Year())
	fmt.Println("Mas 20 dias")
	fecha1 := ahora.Add(time.Hour * 24 * 20)
	fmt.Printf("%v/%v/%v \n", fecha1.Day(), fecha1.Month(), fecha1.Year())
	fmt.Println("Menos 20 dias")
	fecha2 := ahora.Add((time.Hour * 24 * 1) * -20)
	fmt.Printf("%v/%v/%v \n", fecha2.Day(), fecha2.Month(), fecha2.Year())
	fmt.Println("Dentro de un Año")
	fecha3 := ahora.Add(time.Hour * 24 * 365)
	fmt.Printf("%v/%v/%v \n", fecha3.Day(), fecha3.Month(), fecha3.Year())
	fmt.Println("---------Fecha formateada a string-----------")
	fmt.Println(FormatearFecha(ahora))
}

func FormatearFecha(fecha time.Time) string {
	fechaString := fmt.Sprintf("%v/%v/%v \n", fecha.Day(), fecha.Month(), fecha.Year())
	return fechaString
}
