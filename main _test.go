package main_test

import (
	"crud/conect"
	"crud/diversos"
	"crud/menu"
	"fmt"
	"testing"
)

var opt int

func Test_main(t *testing.T) {
	diversos.Limpa()
	for {
		ver := false
		conect.Conect()
		opt = menu.Menu()
		fmt.Println(opt)
		switch opt {
		case 1:
			conect.Create()
			diversos.Limpa()
		case 2:
			conect.ReadDados()
		case 3:
			conect.UpdateUser(menu.MenuUpdate())
			diversos.Limpa()
		case 4:
			var idToDelete int
			fmt.Println("Digite o ID do usuário a ser excluído:")
			fmt.Scan(&idToDelete)
			conect.DeleUser(idToDelete)
			diversos.Limpa()
		case 5:
			ver = true

		}
		if ver == false {
			continue
		} else {
			break
		}
	}

}
