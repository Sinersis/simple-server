package runner

import (
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

func Install() {
	cmd := "curl -L 'https://packages.gitlab.com/install/repositories/runner/gitlab-runner/script.deb.sh' | bash" +
		"apt -y install gitlab-runner;" +
		"gitlab-runner install --user=web --working-directory=/home/web --config=/home/web/.gitlab-runner/config.toml"
	out := exec.Command("bash", "-c", cmd)

	command, err := pty.Start(out)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = io.Copy(os.Stdout, command)
}
