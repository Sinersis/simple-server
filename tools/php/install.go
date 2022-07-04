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
		"apt install php8.1 -y;" +
		"apt -y install php8.1-bcmath;" +
		"apt -y install php8.1-xml;" +
		"apt -y install php8.1-fpm;" +
		"apt -y install php8.1-mysql;" +
		"apt -y install php8.1-zip;" +
		"apt -y install php8.1-intl;" +
		"apt -y install php8.1-ldap;" +
		"apt -y install php8.1-gd;" +
		"apt -y install php8.1-cli;" +
		"apt -y install php8.1-bz2;" +
		"apt -y install php8.1-curl;" +
		"apt -y install php8.1-mbstring;" +
		"apt -y install php8.1-pgsql;" +
		"apt -y install php8.1-opcache;" +
		"apt -y install php8.1-soap;" +
		"apt -y install php8.1-cgi;" +
		"sed -e 's/group = www-data/group = web/g' -e 's/user = www-data/user = web/g' -e 's/listen.owner = www-data/listen.owner = web/g' -e 's/listen.group = www-data/listen.group = web/g' /etc/php/8.1/fpm/pool.d/www.conf > /etc/php/8.1/fpm/pool.d/www.conf;"

	out := exec.Command("bash", "-c", cmd)

	command, err := pty.Start(out)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = io.Copy(os.Stdout, command)
}
