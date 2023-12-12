package diversos

import (
	"os"
	"os/exec"
)

func Limpa() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
