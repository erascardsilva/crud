package menu

import (
	"fmt"
)

func MenuUpdate() (int, string, string, string) {
	var newnome, newemail, newWhatsapp string
	var idnew, decs int

	for {
		fmt.Println("---------------------------------------------------------------------")
		fmt.Println("Atualizar Dados Digite a ID a se atualizar ")
		fmt.Scanln(&idnew)
		fmt.Println("Digite o nome : ")
		fmt.Scanln(&newnome)
		fmt.Println("Digite o email :")
		fmt.Scanln(&newemail)
		fmt.Println("Digite o Whats")
		fmt.Scanln(&newWhatsapp)
		fmt.Printf("Voce digitou : Nome = %v, Email = %v, Whats = %v\n", newnome, newemail, newWhatsapp)
		fmt.Println("Correto? Digite 1 para confirmar, ou qualquer outra tecla para corrigir.")
		fmt.Scanln(&decs)

		if decs == 1 {
			break
		}
	}

	return idnew, newnome, newemail, newWhatsapp
}
