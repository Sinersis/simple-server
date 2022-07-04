package database

import (
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

func Install() {
	cmd := "apt -y install mysql-server;" +
		"apt -y install mysql-client;" +
		"touch /var/run/mysqld/mysqld.sock && chown -R mysql:mysql /var/run/mysqld/mysqld.sock; chmod -R 644 /var/run/mysqld/mysqld.sock; " +
		"touch /var/run/mysqld/mysqld.pid && chown -R mysql:mysql /var/run/mysqld/mysqld.pid;" +
		"usermod -d /var/lib/mysql/ mysql; service mysql start;" + "mysql -u root -e \"ALTER USER 'root'@'localhost' IDENTIFIED BY 'mypassword';\""
	out := exec.Command("bash", "-c", cmd)

	command, err := pty.Start(out)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = io.Copy(os.Stdout, command)
}
