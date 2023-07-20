package main

import (
	"bd/hadlers"
	"bd/modelos"
)

func main() {

	var cliente modelos.Cliente
	cliente.Nombre = "sofia"
	cliente.Correo = "sofia@gmail.com"
	cliente.Telefono = "1212"

	hadlers.Crear(cliente)
	// hadlers.ListaPorId(9)
	hadlers.Listar()
}
