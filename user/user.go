package user

import (
	"easy-server/pwd"
	"easy-server/service"
	"fmt"
	"github.com/creack/pty"
	"io"
	"log"
	"os"
	"os/exec"
)

func AddExecutor(enableStdOur bool) {

	if !enableStdOur {
		fmt.Print(service.PrintInfoText("Executor User"))
	}

	cmd := "useradd -m -p s3cr3tP4ssW0rd -s /bin/bash web; usermod -aG web web"
	out := exec.Command("bash", "-c", cmd)

	if enableStdOur {

		command, err := pty.Start(out)

		if err != nil {
			fmt.Print(service.PrintErrorText())
			log.Fatal(err)
		}

		_, _ = io.Copy(os.Stdout, command)

		pwd.AddPwdToFile("web", "s3cr3tP4ssW0rd")
	} else {
		err := out.Run()

		if err != nil {
			fmt.Print(service.PrintErrorText())
			log.Fatal(err)
		}

		pwd.AddPwdToFile("web", "s3cr3tP4ssW0rd")
		fmt.Print(service.PrintDoneText())
	}

}

func AddSudo(enableStdOur bool) {

	if !enableStdOur {
		fmt.Print(service.PrintInfoText("Sudo User"))
	}

	cmd := "useradd -m -p s3cr3tP4ssW0rd -s /bin/bash service; usermod -aG sudo service"
	out := exec.Command("bash", "-c", cmd)

	if enableStdOur {
		command, err := pty.Start(out)

		if err != nil {
			fmt.Print(service.PrintErrorText())
			log.Fatal(err)
		}

		_, _ = io.Copy(os.Stdout, command)

		pwd.AddPwdToFile("service", "s3cr3tP4ssW0rd")
		disableSudoPassword("service")
	} else {

		err := out.Run()

		if err != nil {
			fmt.Print(service.PrintErrorText())
		}

		pwd.AddPwdToFile("service", "s3cr3tP4ssW0rd")
		disableSudoPassword("service")

		fmt.Print(service.PrintDoneText())
	}

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
