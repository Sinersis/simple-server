package main

import (
	"github.com/codegangsta/cli"
	"github.com/creack/pty"
	"io"
	"os"
	"os/exec"
)

/**
Some config for clear server
Execute from root or sudo user
*/
func main() {

	app := cli.NewApp()
	app.Name = "Clear install Server"
	app.Usage = "Install all you need for new php web server"
	app.Commands = []cli.Command{
		{
			Name:   "check",
			Usage:  "Check Command",
			Action: checkCommand,
		},
	}

	app.Run(os.Args)
}

func checkCommand(ctx *cli.Context) {

	/**
	Best Practice with commands
	*/

	c := exec.Command("ls")
	f, err := pty.Start(c)
	if err != nil {
		panic(err)
	}

	go func() {
		f.Write([]byte{}) // EOT
	}()
	io.Copy(os.Stdout, f)

	//cmd := "brew upgrade"
	//out := exec.Command("bash", "-c", cmd)
	//
	//pr, err := out.CombinedOutput()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(string(pr))
}

func checkAnyUsers(ctx *cli.Context) {
	//checked if any users with name service or web
}

func baseCommandsBeforeInstall(ctx *cli.Context) {
	// apt update
	// apt upgrade
}

func createSudoUser(ctx *cli.Context) {
	// create non root sudo user
}

func createExecuteUser(ctx *cli.Context) {
	// create web or execute user
}

func installNode(ctx *cli.Context) {
	// install node (before last version) (after checked version or nvm)
}

func installNginx(ctx *cli.Context) {
	// install Nginx (last version) (for execute user config)
}
func installPhp(ctx *cli.Context) {
	// install Php (**with base library) (before last version) (after checked version)
}
func installMySql(ctx *cli.Context) {
	// install MySql (last version)
}

func installGitlabRunner(ctx *cli.Context) {
	// install gitlab runner (for execute user config)
}
