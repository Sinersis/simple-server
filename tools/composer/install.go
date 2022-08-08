package composer

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
		fmt.Print(service.PrintInfoText("Composer"))
	}

	cmd := "curl -sS https://getcomposer.org/installer -o /tmp/composer-setup.php; php /tmp/composer-setup.php --install-dir=/usr/local/bin --filename=composer"
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
