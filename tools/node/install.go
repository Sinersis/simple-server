package node

import (
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

func Install() {
	cmd := "wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash;"
	out := exec.Command("bash", "-c", cmd)

	command, err := pty.Start(out)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = io.Copy(os.Stdout, command)
}
