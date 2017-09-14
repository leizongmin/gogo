package cmd

import (
	"fmt"
	"log"
)

// Run 执行其他命令
func Run(args []string) {
	pkg, exec := getPackageInfoAndExec(false)
	if !isWorkspaceDirExists(pkg.Dir.Pwd) {
		log.Println(`"_workspace" directory doesn't exists, please run "gogo init" before.`)
		return
	}
	exec.setDir(pkg.Dir.PwdUnderWorkspace)
	if len(args) >= 1 {
		ret := exec.run(args[0], args[1:]...)
		fmt.Print(ret)
	} else {
		RunHelp(args)
	}
}

// RunHelp 命令帮助
func RunHelp(args []string) {
	fmt.Println(`
usage: gogo run dep init
       gogo run env

run any command.
	`)
}
