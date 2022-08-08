package database

import (
	"easy-server/service"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

func Install(enableStdOur bool) {

	if !enableStdOur {
		fmt.Print(service.PrintInfoText("Database"))
	}

	cmd := "apt -y install mysql-server;" +
		"apt -y install mysql-client;" +
		"touch /var/run/mysqld/mysqld.sock && chown -R mysql:mysql /var/run/mysqld/mysqld.sock; chmod -R 644 /var/run/mysqld/mysqld.sock; " +
		"touch /var/run/mysqld/mysqld.pid && chown -R mysql:mysql /var/run/mysqld/mysqld.pid;" +
		"usermod -d /var/lib/mysql/ mysql; service mysql start;" + "mysql -u root -e \"ALTER USER 'root'@'localhost' IDENTIFIED BY 'mypassword';\""
	out := exec.Command("bash", "-c", cmd)

	if enableStdOur {
		command, err := pty.Start(out)

		if err != nil {
			fmt.Print(service.PrintErrorText())
			log.Fatal(err)
		}

		_, _ = io.Copy(os.Stdout, command)
	} else {
		err := out.Run()

		if err != nil {
			fmt.Print(service.PrintErrorText())
		}

		fmt.Print(service.PrintDoneText())
	}
}
