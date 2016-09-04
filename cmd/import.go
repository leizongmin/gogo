package cmd

import "fmt"

func Import(args []string) {

	_, exec := getPackageInfoAndExec(true)

	// `go get ${argv.join(' ')}`

	newArgs := combineStringArray([]string{"get"}, args)
	exec("go", newArgs...)

	fmt.Println("\nOK")

}

func ImportHelp(args []string) {
	fmt.Println(`
usage: gogo import package1 package2 package3

add import packages and save to package.yaml file
	`)
}
