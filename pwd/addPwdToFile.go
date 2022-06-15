package pwd

import (
	"github.com/creack/pty"
	"io"
	"log"
	"os"
	"os/exec"
)

func AddPwdToFile(userName string, userPwd string) {
	cmd := "echo 'username: " + userName + " password: " + userPwd + "' >> ./password"
	out := exec.Command("bash", "-c", cmd)

	command, err := pty.Start(out)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = io.Copy(os.Stdout, command)
}
