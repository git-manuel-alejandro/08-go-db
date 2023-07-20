package conectar

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

var Db *sql.DB

func Conectar() {
	errorVariables := godotenv.Load()

	if errorVariables != nil {
		panic(errorVariables)
	}

	connection, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_SERVER")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}

	Db = connection

}

func CerrarConexion() {
	Db.Close()
}
