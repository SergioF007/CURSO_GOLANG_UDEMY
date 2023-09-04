package main

import (
	"log"
	"os"
)

// logs

func main() {
	//err := errors.New("Este es un error fatal de prueba")
	//log.Fatal(err) // esto me detine la ejecuion
	//log.Println("Straing the application ...")
	//log.Fatalln("Fatal: Fatal error signal")
	//log.Fatalf("Fatal: Fatal error signal", "Var err")
	//log.Panic("Error: Fatal error message")
	//log.Panicln("Error: Panic error message")
	//error := "todo esta mal"
	//log.Panicf("error %s", error)
	f, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	// cerrar el archivo cuando se termina la ejecucion
	defer f.Close()
	log.SetOutput(f)
	log.Println("Error: ", err)

}

/*
// modulo os | argumentos
func main() {

	// go run main.go -nombre Sergio -edad 22
	nombre := flag.String("nombre", "", "El nombre de la persona")
	edad := flag.Int("edad", 27, "La edad de la persona")

	flag.Parse()

	fmt.Printf("Tu nombre es: %s, tines %d años\n", *nombre, *edad)
}
*/

/*
// math/rand
func main() {
	numAleatorio := rand.Intn(100)
	fmt.Println(numAleatorio)
	min := 100
	max := 1000
	rand.Seed(time.Now().UnixNano())
	numAleatorio2 := rand.Intn(max-min) + min
	fmt.Println(numAleatorio2)
	fmt.Println(generatePassword(20, 1, 1, 1))

}

var (
	lowerCharSet   = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!#$%&*+@"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func generatePassword(passwordLength, minSpecialChar, miniNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		password.WriteString(string(specialCharSet[rand.Intn(len(specialCharSet))]))
	}

	// Set numeric
	for i := 0; i < miniNum; i++ {
		password.WriteString(string(numberSet[rand.Intn(len(numberSet))]))
	}

	// Set uppercase
	for i := 0; i < minUpperCase; i++ {
		password.WriteString(string(upperCharSet[rand.Intn(len(upperCharSet))]))
	}

	/*
		    // Set lowercase
			for i := 0; i < passwordLength-miniNum-minUpperChar-minSpecialChar; i++ {
		        password.WriteString(string(lowerCharSet[rand.Intn(len(lowerCharSet))]))
		    }

*/
/*
	remainingLength := passwordLength - minSpecialChar - miniNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		password.WriteString(string(allCharSet[rand.Intn(len(allCharSet))]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

*/

/*
// modulo strings
func main() {
	cadena := "mi muñeca me hablo"
	fmt.Println(strings.ToUpper(cadena))
	fmt.Println(strings.ToLower(cadena))
	fmt.Println("-----------------Se para la cadena por cada una de las letras---------------------")
	letras := strings.Split(cadena, "")
	fmt.Println(letras)
	fmt.Println("--------Buscar una palabra en un cadena de texto--------")
	pos := strings.Index(cadena, "muñeca")
	if pos == -1 {
		fmt.Println("No se encontro la palabra")
	} else {
		fmt.Println("La palabra se encuentra en la cadena de texto", cadena, " a partir de la posicion", pos)
	}
	fmt.Println("-------------Repetir la cadena de texto------------")
	repetida := strings.Repeat(cadena, 2)
	fmt.Println(repetida)
	fmt.Println("----------------Remplazar una palabra de la cadena---------------- ")
	cadenaNueva := strings.Replace(cadena, "mi", "la", -1)
	fmt.Println(cadenaNueva)
	fmt.Println("----------------Imprimer una seccion de la nueva cadena---------------- ")
	fmt.Println(cadenaNueva[0:10])

}
*/

/*
// modulo time
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

*/
