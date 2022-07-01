package php

import (
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

func Install() {

	cmd := "add-apt-repository ppa:ondrej/php -y;" +
		"apt install php8.1;" +
		"apt install php8.1-{bcmath,xml,fpm,mysql,zip,intl,ldap,gd,cli,bz2,curl,mbstring,pgsql,opcache,soap,cgi};" +
		"sed -e 's/group = www-data/group = web/g' -e 's/user = www-data/user = web/g' -e 's/listen.owner = www-data/listen.owner = web/g' -e 's/listen.group = www-data/listen.group = web/g' www.conf > www1.conf;"

	out := exec.Command("bash", "-c", cmd)

	command, err := pty.Start(out)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = io.Copy(os.Stdout, command)
}
