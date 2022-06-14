package nginx

import (
	"github.com/creack/pty"
	"io"
	"log"
	"os"
	"os/exec"
)

func Install() {
	cmd := "apt -y install nginx; echo 'user web;' >> /etc/nginx/nginx.conf"
	out := exec.Command("bash", "-c", cmd)

	command, err := pty.Start(out)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = io.Copy(os.Stdout, command)
}
