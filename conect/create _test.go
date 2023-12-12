package conect

import (
	"crud/menu"
	"fmt"
	"log"
	"testing"
)

func Test_Create(t *testing.T) {
	db := Conect()

	defer db.Close()
	Conect()
	dataCreate := menu.MenuCreate()
	fmt.Println(dataCreate)
	result, err := db.Exec(dataCreate)
	if err != nil {
		log.Fatal(err)
	}
	userID, _ := result.LastInsertId()

	fmt.Printf("Criado usuario com id %v", userID)

}
