package menu

import "fmt"

func MenuCreate() string {
	var name, email, whats string

	fmt.Println("Digite o nome do usuário:")
	fmt.Scanln(&name)
	fmt.Println("Digite o email:")
	fmt.Scanln(&email)
	fmt.Println("Digite o Whats:")
	fmt.Scanln(&whats)

	dataCreate := fmt.Sprintf("INSERT INTO usuarios (Nome, Email, Whats) VALUES ('%s', '%s', %s)", name, email, whats)

	return dataCreate
}
