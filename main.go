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

/**
Some config for clear server
Execute from root or sudo user
*/
func main() {

	app := cli.NewApp()
	app.Name = "Easy Server Config"
	app.Usage = "This is experimental program. Working with only Ubuntu Server. This tool installed all for you need for your PHP server, for framework Laravel"
	app.Commands = []cli.Command{
		{
			Name:   "start",
			Usage:  "Base Install for clear server",
			Action: run,
		},
	}

	_ = app.Run(os.Args)
}
func run(_ *cli.Context) {
	user.AddExecutor()
	user.AddSudo()

	preinstall.PreInstall()
	git.Install()
	nginx.Install()
	node.Install()
	php.Install()
	composer.Install()
	runner.Install()
	database.Install()

}
