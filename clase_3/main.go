package main

import (
	"github.com/SergioF007/cursoGoUdemy/handlers"
)

func main() {
	//handlers.Listar()
	//handlers.ListarPorId(2)

	/*
		fmt.Println("-----------------Crear Cliente-----------------------")
		var cliente models.Cliente
		cliente.Nombre = "Jose Ignacio Fernández"
		cliente.Correo = "jsf@cursogo.com"
		cliente.Telefono = "999999999"

		handlers.Insert(cliente)
		handlers.Listar()

		fmt.Println("------------------Editar Cliente---------------------")

		cliente2 := models.Cliente{Nombre: "Jose Ruiz", Correo: "jsf2@cursogo.com", Telefono: "88888888"}
		handlers.Editar(cliente2, 4)
		handlers.Listar()

		fmt.Println("-----------------Eliminar Cliente--------------------")
		handlers.Eliminar(6)
		handlers.Listar()
	*/

	handlers.Ejecutar()

}
