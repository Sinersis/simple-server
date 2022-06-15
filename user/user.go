package user

import (
	"easy-server/pwd"
	"github.com/creack/pty"
	"io"
	"log"
	"os"
	"os/exec"
)

func AddExecutor() {
	cmd := "useradd -m -p s3cr3tP4ssW0rd -s /bin/bash web; usermod -aG web web"
	out := exec.Command("bash", "-c", cmd)

	command, err := pty.Start(out)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = io.Copy(os.Stdout, command)
	pwd.AddPwdToFile("web", "s3cr3tP4ssW0rd")
}

func AddSudo() {
	cmd := "useradd -m -p s3cr3tP4ssW0rd -s /bin/bash service; usermod -aG sudo service"
	out := exec.Command("bash", "-c", cmd)

	command, err := pty.Start(out)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = io.Copy(os.Stdout, command)
	pwd.AddPwdToFile("service", "s3cr3tP4ssW0rd")
	disableSudoPassword("service")
}

func disableSudoPassword(userName string) {

	cmd := "echo 'service ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers.d/" + userName

	out := exec.Command("bash", "-c", cmd)

	command, err := pty.Start(out)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = io.Copy(os.Stdout, command)

}
