package main

import (
	"bd/hadlers"
	"bd/modelos"
)

func main() {

	var cliente modelos.Cliente
	cliente.Nombre = "mnolo"
	cliente.Correo = "mnolo@gmail.com"
	cliente.Telefono = "noloa"

	hadlers.Editar(cliente, 10)
	// hadlers.Crear(cliente)
	// hadlers.ListaPorId(9)
	hadlers.Listar()
}
