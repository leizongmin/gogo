package cmd

import "fmt"

// Go 执行 go 命令
func Go(args []string) {
	pkg, exec := getPackageInfoAndExec(false)
	if !isWorkspaceDirExists(pkg.Dir.Pwd) {
		fmt.Println(`
"_workspace" directory doesn't exists, please run "gogo init" before.
		`)
		return
	}
	exec.setDir(pkg.Dir.Pwd)
	exec.run("go", args...)
}

// GoHelp 命令帮助
func GoHelp(args []string) {
	fmt.Println(`
usage: gogo - env
       gogo - vet
       gogo - get package1 package2

run any go command.
	`)
}
