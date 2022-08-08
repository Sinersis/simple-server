package runner

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
		fmt.Print(service.PrintInfoText("Gitlab Runner"))
	}

	cmd := "curl -L 'https://packages.gitlab.com/install/repositories/runner/gitlab-runner/script.deb.sh' | bash;" +
		"apt -y install gitlab-runner;" +
		"gitlab-runner uninstall;" +
		"gitlab-runner install --user=web --working-directory=/home/web --config=/home/web/.gitlab-runner/config.toml"
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
			log.Fatal(err)
		}

		fmt.Print(service.PrintDoneText())
	}
}
