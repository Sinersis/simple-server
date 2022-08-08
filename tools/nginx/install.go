package nginx

import (
	"easy-server/service"
	"fmt"
	"github.com/creack/pty"
	"io"
	"log"
	"os"
	"os/exec"
)

func Install(enableStdOur bool) {

	if !enableStdOur {
		fmt.Print(service.PrintInfoText("Nginx"))
	}

	cmd := "apt -y install nginx; echo 'user web;' >> /etc/nginx/nginx.conf"
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
