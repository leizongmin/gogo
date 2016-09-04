package cmd

import "fmt"

func Go(args []string) {

	pkg, exec := getPackageInfoAndExec(false)

	if !isWorkspaceDirExists(pkg.Dir.Pwd) {
		fmt.Println(`
"_workspace" directory doesn't exists, please run "gogo init" before.
		`)
		return
	}

	exec("go", args...)

}

func GoHelp(args []string) {
	fmt.Println(`
usage: gogo - env
       gogo - vet
       gogo - get package1 package2

run any go command.
	`)
}
