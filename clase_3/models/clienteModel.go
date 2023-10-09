package models

type Cliente struct {
	Id       int
	Nombre   string
	Correo   string
	Telefono string
}

// struct para crear una lista de cliente
type Clientes []Cliente
