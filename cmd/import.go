package cmd

import "fmt"

func Import(args []string) {

	_, exec := getPackageInfoAndExec(true)

	// `go get ${argv.join(' ')}`

	newArgs := make([]string, len(args)+1)
	newArgs[0] = "get"
	copy(newArgs[1:], args)

	exec("go", newArgs...)

	fmt.Println("\nOK")

}
