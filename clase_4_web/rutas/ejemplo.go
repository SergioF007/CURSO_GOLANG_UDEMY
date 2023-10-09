package rutas

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/ejemplo/home.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(w, nil)
	}

}

func Nosotros(w http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/ejemplo/nosotros.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(w, nil)
	}
}

func Parametros(w http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/ejemplo/parametros.html")
	vars := mux.Vars(req)
	texto := "le estamos pasando texto desde el metodo en go"
	data := map[string]string{
		"id":    vars["id"], // asi tenemos que llamar a los valores e el html para poder tomar su respuesta
		"slug":  vars["slug"],
		"texto": texto,
	}
	if err != nil {
		panic(err)
	} else {
		template.Execute(w, data)
	}

}

func ParametrosQueryString(w http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/ejemplo/parametros_query_string.html")
	data := map[string]string{
		"id":   req.URL.Query().Get("id"),
		"slug": req.URL.Query().Get("slug"),
	}
	if err != nil {
		panic(err)
	} else {
		template.Execute(w, data)
	}

}

type Habilidad struct {
	Nombre string
}

type Datos struct {
	Nombre      string
	Edad        int
	Perfil      int
	Habilidades []Habilidad
}

func Estructuras(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles("templates/ejemplo/estructuras.html")
	if err != nil {
		panic(err)
	}
	hablidad1 := Habilidad{"Inteligencia"}
	hablidad2 := Habilidad{"VideoJuegos"}
	hablidad3 := Habilidad{"Programación"}

	Habilidades := []Habilidad{hablidad1, hablidad2, hablidad3}
	template.Execute(w, Datos{"Sergio", 27, 1, Habilidades})
	/*
		template, err := template.ParseFiles("templates/ejemplo/estructuras.html")
		if err != nil {
			panic(err)
		} else {
			template.Execute(w, Datos{"Sergio", 27, 1})
		}
	*/
}

/*
func Home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hola Mundo desde Golang")

}
func Nosotros(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Pagina sobre nosotros prueba con fresh")
}
func Parametros(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	fmt.Fprintln(w, "ID = "+vars["id"]+" | SLUG =  "+vars["slug"])

}
func ParametrosQueryString(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, req.URL)
	fmt.Fprintln(w, req.URL.RawQuery)            // metoma los parametros queryString
	fmt.Fprintln(w, req.URL.Query())             // metoma los parametros en un mapa
	fmt.Fprintln(w, req.URL.Query().Get("id"))   // metoma el valor del parametro de id
	fmt.Fprintln(w, req.URL.Query().Get("slug")) // metoma el valor del parametro de slug

	id := req.URL.Query().Get("id")
	slug := req.URL.Query().Get("slug")
	fmt.Fprintln(w, "--------------------------------")
	fmt.Fprintf(w, " id: %s | slug: %s ", id, slug)

}

*/
