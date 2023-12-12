package conect

import (
	"fmt"
	"log"
	"testing"
)

func Test_DeleUser(t *testing.T) {
	id := 1
	db := Conect()
	defer db.Close()

	var userexist bool
	err := db.QueryRow("SELECT 1 FROM usuarios WHERE id = ?", id).Scan(&userexist)
	if err != nil {
		log.Fatal(err)
	}
	if !userexist {
		fmt.Println("Usuario nao existe")
		return
	}
	_, err = db.Exec("DELETE FROM usuarios WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Usuario Excluido !!")
}
