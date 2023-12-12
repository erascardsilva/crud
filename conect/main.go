package conect

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "./data/banco.db"

func Conect() *sql.DB {
	if err := os.MkdirAll(filepath.Dir(dbName), os.ModePerm); err != nil {
		panic(err)
	}

	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		createDatabase()
	}

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("CONECTADO")
	}

	_, err = db.Exec(`
CREATE TABLE IF NOT EXISTS usuarios (
	Id INTEGER PRIMARY KEY AUTOINCREMENT,
	Nome TEXT,
	Email TEXT,
	Whats INTEGER
)
`)
	if err != nil {
		panic(err)
	}

	return db
}

func createDatabase() {
	file, err := os.Create(dbName)
	if err != nil {
		panic(err)
	}
	file.Close()
	fmt.Println("Banco de dados criado em", dbName)
}
