package diversos

import (
	"os"
	"os/exec"
	"testing"
)

func Test_Limpa(t *testing.T) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
