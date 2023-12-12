package menu

import (
	"fmt"
)

func Menu() int {
	var opt int
	for {
		fmt.Println("=====================================")
		fmt.Println("Menu Banco de Daados")
		fmt.Println("Escolha uma opção do CRUD")
		fmt.Println("1 para incluir um usuario")
		fmt.Println("2 para verificar os usuarios")
		fmt.Println("3 para atualizar usuario")
		fmt.Println("4 para apagar um usuario")
		fmt.Println("5 para sair")
		fmt.Println("=====================================")

		fmt.Println("Digite uma opção")
		fmt.Scan(&opt)
		if opt <= 0 || opt >= 6 {
			fmt.Println("Tente novamente apenas de 1 a 5")
			continue

		} else {
			break
		}
	}
	return opt
}
