package cmd

import "fmt"

func Go(args []string) {

	_, exec := getPackageInfoAndExec(false)

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
