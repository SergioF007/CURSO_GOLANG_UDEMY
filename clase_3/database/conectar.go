package database

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func Conectar() {

	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}
	conection, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_SERVER")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}
	DB = conection

}

func CerrarConexion() {
	DB.Close()
}
