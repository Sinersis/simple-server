package preinstall

import (
	"github.com/creack/pty"
	"io"
	"log"
	"os"
	"os/exec"
)

func PreInstall() {
	cmd := "apt update && apt -y upgrade "
	out := exec.Command("bash", "-c", cmd)

	command, err := pty.Start(out)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = io.Copy(os.Stdout, command)
}
