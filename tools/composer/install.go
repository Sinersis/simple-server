package composer

import (
	"github.com/creack/pty"
	"io"
	"log"
	"os"
	"os/exec"
)

func Install() {
	//cmd := "apt -y install composer"
	cmd := "curl -sS https://getcomposer.org/installer -o /tmp/composer-setup.php; php /tmp/composer-setup.php --install-dir=/usr/local/bin --filename=composer"
	out := exec.Command("bash", "-c", cmd)

	command, err := pty.Start(out)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = io.Copy(os.Stdout, command)
}
