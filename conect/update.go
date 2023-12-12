package conect

import (
	"fmt"
	"log"
)

func UpdateUser(idnew int, newnome, newemail, newWhatsapp string) {
	db := Conect()
	defer db.Close()

	var userexist bool
	err := db.QueryRow("SELECT 1 FROM usuarios WHERE id = ?", idnew).Scan(&userexist)
	if err != nil {
		log.Fatal(err)
	}
	if !userexist {
		fmt.Println("Usuario não encontrado")
		return
	}

	_, err = db.Exec("UPDATE usuarios SET nome = ?, email = ? , whats = ? WHERE id = ?", newnome, newemail, newWhatsapp, idnew)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Dados atualizados")
}
