package node

import (
	"easy-server/service"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

func Install(enableStdOur bool) {

	if !enableStdOur {
		fmt.Print(service.PrintInfoText("Node"))
	}

	cmd := "wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash;"
	out := exec.Command("bash", "-c", cmd)

	if enableStdOur {
		command, err := pty.Start(out)

		if err != nil {
			fmt.Print(service.PrintErrorText())
			log.Fatal(err)
		}
		_, _ = io.Copy(os.Stdout, command)
	} else {
		err := out.Run()

		if err != nil {
			fmt.Print(service.PrintErrorText())
		}

		fmt.Print(service.PrintDoneText())
	}
}
