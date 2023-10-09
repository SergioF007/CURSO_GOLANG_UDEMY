package handlers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	conectar "github.com/SergioF007/cursoGoUdemy/database"
	"github.com/SergioF007/cursoGoUdemy/models"
)

func Listar() {
	conectar.Conectar()
	sql := "SELECT id, nombre, correo, telefono FROM clientes ORDER BY id DESC;"
	resultQuery, err := conectar.DB.Query(sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conectar.CerrarConexion()
	/*
		clientes := models.Clientes{}
		for resultQuery.Next() {
			cliente := models.Cliente{}
			resultQuery.Scan(&cliente.Id, &cliente.Nombre, &cliente.Correo, &cliente.Telefono)
			clientes = append(clientes, cliente)
		}
		fmt.Println(clientes)
	*/

	fmt.Println("\nListado de Clientes")
	fmt.Println("----------------------------------------------------------------")
	for resultQuery.Next() {

		var cliente = models.Cliente{}
		err := resultQuery.Scan(&cliente.Id, &cliente.Nombre, &cliente.Correo, &cliente.Telefono)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Id: %v | Nombre: %v | Correo: %s | Telefono: %s \n", cliente.Id, cliente.Nombre, cliente.Correo, cliente.Telefono)
		fmt.Println("----------------------------------------------------")
	}

}

func ListarPorId(id int) {
	conectar.Conectar()
	sql := "SELECT id, nombre, correo, telefono FROM clientes WHERE id=?;"
	resultQuery, err := conectar.DB.Query(sql, id)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conectar.CerrarConexion()
	fmt.Println("\nListado de Clientes por Id")
	fmt.Println("----------------------------------------------------------------")
	for resultQuery.Next() {

		var cliente = models.Cliente{}
		err := resultQuery.Scan(&cliente.Id, &cliente.Nombre, &cliente.Correo, &cliente.Telefono)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Id: %v | Nombre: %v | Correo: %s | Telefono: %s \n", cliente.Id, cliente.Nombre, cliente.Correo, cliente.Telefono)
		fmt.Println("----------------------------------------------------")
	}
}

func Insert(cliente models.Cliente) {
	conectar.Conectar()
	sql := "INSERT INTO clientes VALUES (null, ?,?,?);"
	result, err := conectar.DB.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
	fmt.Println("Se creo el registro exitosamente")
}

func Editar(cliente models.Cliente, id int) {

	conectar.Conectar()
	sql := "UPDATE clientes SET nombre=?, correo=?, telefono=? WHERE id=?"
	result, err := conectar.DB.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono, id)
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println(result)
	fmt.Println("Se edito el registro exitosamente")
}

func Eliminar(id int) {
	conectar.Conectar()
	sql := "DELETE FROM clientes WHERE id=?;"
	result, err := conectar.DB.Exec(sql, id)
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println(result)
	fmt.Println("Se elimino el registro exitosamente")

}

// ################################# Funciones de trabajo
var ID int
var nombre, correo, telefono string

func Ejecutar() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Bienvenido al sistema de clientes")
	fmt.Print("Selecciones una opcion: \n\n")
	fmt.Print("1. Listar Clientes \n")
	fmt.Print("2. Listar Cliente por ID \n")
	fmt.Print("3. Crear Cliente \n")
	fmt.Print("4. Editar Cliente \n")
	fmt.Print("5. Eliminar Cliente \n")
	if scanner.Scan() {

		for {

			if scanner.Text() == "1" {
				Listar()
				return
			}
			if scanner.Text() == "2" {
				fmt.Println("Ingrese el ID del cliente que desea consultar")
				fmt.Scanln(&ID)
				ListarPorId(ID)
				return
			}
			if scanner.Text() == "3" {
				scanner := bufio.NewScanner(os.Stdin)

				fmt.Println("Ingrese el nombre del cliente")
				scanner.Scan()
				nombre := scanner.Text()
				fmt.Println("Ingrese el correo del cliente")
				scanner.Scan()
				correo := scanner.Text()
				fmt.Println("Ingrese el telefono del cliente")
				scanner.Scan()
				telefono := scanner.Text()

				Insert(models.Cliente{Nombre: nombre, Correo: correo, Telefono: telefono})
				return
			}
			if scanner.Text() == "4" {
				fmt.Println("Ingrese el ID del cliente que desea editar")
				scanner.Scan()
				ID, _ := strconv.Atoi(scanner.Text())
				fmt.Println("Ingrese el nombre del cliente")
				scanner.Scan()
				nombre := scanner.Text()
				fmt.Println("Ingrese el correo del cliente")
				scanner.Scan()
				correo := scanner.Text()
				fmt.Println("Ingrese el telefono del cliente")
				scanner.Scan()
				telefono := scanner.Text()
				Editar(models.Cliente{Nombre: nombre, Correo: correo, Telefono: telefono}, ID)
				return
			}
			if scanner.Text() == "5" {
				fmt.Println("Ingrese el ID del cliente que desea eliminar")
				fmt.Scanln(&ID)
				Eliminar(ID)
				return
			}

		}
	}

}
