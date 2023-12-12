package menu

import (
	"fmt"
)

func Test_MenuCreate() string {
	var name, email, whats string

	fmt.Println("Digite o nome do usuário:")
	//fmt.Scanln(&name)
	name = "TESTE"
	fmt.Println("Digite o email:")
	//fmt.Scanln(&email)
	email = "email@test"
	fmt.Println("Digite o Whats:")
	//fmt.Scanln(&whats)
	whats = "111111111111111"
	dataCreate := fmt.Sprintf("INSERT INTO usuarios (Nome, Email, Whats) VALUES ('%s', '%s', %s)", name, email, whats)

	return dataCreate
}
