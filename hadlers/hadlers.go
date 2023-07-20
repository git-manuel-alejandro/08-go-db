package hadlers

import (
	"bd/conectar"
	"bd/modelos"
	"fmt"
)

func Listar() {
	conectar.Conectar()
	sql := "select id, nombre, correo, telefono from clientes;"
	datos, err := conectar.Db.Query(sql)

	if err != nil {
		fmt.Println(err)
	}
	defer conectar.CerrarConexion()

	clientes := modelos.Clientes{}

	for datos.Next() {
		dato := modelos.Cliente{}
		datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		clientes = append(clientes, dato)
	}

	fmt.Println(clientes)
}

func ListaPorId(id int) {

	conectar.Conectar()
	sql := "select id, nombre, correo, telefono from clientes where id=?;"
	datos, err := conectar.Db.Query(sql, id)

	if err != nil {
		fmt.Println(err)
	}
	defer conectar.CerrarConexion()

	for datos.Next() {
		var dato modelos.Cliente
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(dato)
	}

}

func Crear(cliente modelos.Cliente) {
	conectar.Conectar()
	sql := "insert into clientes values(null, ?,?,?,?);"
	result, err := conectar.Db.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono, "")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
	fmt.Println("creado el registro")

}
