package main

import (
	"clase_4_web/rutas"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	mux := mux.NewRouter()
	// rutas
	mux.HandleFunc("/", rutas.Home)
	mux.HandleFunc("/nosotros", rutas.Nosotros)
	mux.HandleFunc("/parametros/{id:.*}/{slug:.*}", rutas.Parametros)
	// como procesar parametros en formatos queryString
	mux.HandleFunc("/parametros-querystring", rutas.ParametrosQueryString)
	// Estructuras
	mux.HandleFunc("/estructuras", rutas.Estructuras)

	// Configuracion para que la app reconozca cualquier recurso estatico
	// Archivos estaticos hacia mux
	s := http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))
	mux.PathPrefix("/public/").Handler(s)

	//ejecucuin del servidor
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}

	server := &http.Server{
		Addr:         "127.0.0.1:" + os.Getenv("DB_PORT"),
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Corriendo Servidor desde: http://127.0.0.1:" + os.Getenv("DB_PORT"))
	log.Fatal(server.ListenAndServe())

}

/*
func main() {
	//mux := http.NewServeMux()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hola mundo") //Fprintf me permite imprimir en pantalla dentro del navegador
	})

	fmt.Println("Corriendo Servidor desde: localhost:8081")
	log.Fatal(http.ListenAndServe("localhost:8081", nil))

}
*/
