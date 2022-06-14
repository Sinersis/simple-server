package main

import (
	"easy-server/tools/git"
	"easy-server/tools/nginx"
	"easy-server/tools/node"
	"easy-server/tools/php"
	"easy-server/tools/preinstall"
	"easy-server/tools/runner"
	"easy-server/user"
	"github.com/codegangsta/cli"
	"os"
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
			Name:   "start",
			Usage:  "Test run",
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
	//composer.Install()
	runner.Install()

}

//func installNode(ctx *cli.Context) {
//	preinstall.PreInstall()
//}
//
//func installNginx(ctx *cli.Context) {
//	preinstall.PreInstall()
//}
//func installPhp(ctx *cli.Context) {
//	preinstall.PreInstall()
//}
//func installMySql(ctx *cli.Context) {
//	preinstall.PreInstall()
//}
//
//func installGitlabRunner(ctx *cli.Context) {
//	preinstall.PreInstall()
//}
//
//func installGit(ctx *cli.Context) {
//	preinstall.PreInstall()
//}
//
//func installComposer(ctx *cli.Context) {
//	preinstall.PreInstall()
//}
