package php

import (
	"github.com/creack/pty"
	"io"
	"log"
	"os"
	"os/exec"
)

func Install() {
	cmd := "apt -y install php php-fpm php-bcmath php-ctype php-curl php-dom php-fileinfo php-json php-mbstring" +
		" php-pdo php-tokenizer php-mysql php-xml"

	out := exec.Command("bash", "-c", cmd)

	command, err := pty.Start(out)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = io.Copy(os.Stdout, command)
}
