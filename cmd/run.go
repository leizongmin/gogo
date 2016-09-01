package cmd

import "strings"

func Run(args []string) {

	_, exec := getPackageInfoAndExec(true)

	// `go get ${argv.join(' ')}`

	exec("go", strings.Join(args, " "))

}
