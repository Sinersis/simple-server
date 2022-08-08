package main

import (
	"easy-server/tools/composer"
	"easy-server/tools/database"
	"easy-server/tools/git"
	"easy-server/tools/nginx"
	"easy-server/tools/node"
	"easy-server/tools/php"
	"easy-server/tools/preinstall"
	"easy-server/tools/runner"
	"easy-server/user"
	"os"

	"github.com/codegangsta/cli"
)

type Author struct {
	Name  string
	Email string
}

/**
Some config for clear server
Execute from root or sudo user
*/
func main() {

	app := cli.NewApp()
	app.Name = "Easy Server Config"
	app.Usage = "This is experimental program. Working with only Ubuntu Server. This tool installed all for you need for your PHP server, for framework Laravel"
	app.Version = "1.0.0"
	app.Authors = []cli.Author{
		cli.Author{Name: "Andrey Melentev", Email: "melentev.av@yahoo.com"},
	}
	app.Commands = []cli.Command{
		{
			Name:  "default",
			Usage: "Install default tools for web project on PHP",
			Action: func(ctx *cli.Context) {
				user.AddExecutor(false)
				user.AddSudo(false)

				preinstall.PreInstall(false)
				git.Install(false)
				nginx.Install(false)
				php.Install(false)
				composer.Install(false)
				database.Install(false)
				runner.Install(false)
				node.Install(false)
			},
		},
	}

	_ = app.Run(os.Args)
}
