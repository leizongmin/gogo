package cmd

import (
	"fmt"
)

// Run 执行其他命令
func Run(args []string) {
	pkg, exec := getPackageInfoAndExecAndEnsureInited(false)
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
Usage: gogo run dep init
       gogo run env

run any command.
	`)
}
