package cmd

import (
	"fmt"
)

func Help(args []string) {

	Version(args)

	fmt.Printf(`
Usage:

    gogo build      compiles the current project
    gogo clean      removes object files from current project directory
    gogo help       print the help information
    gogo import     add import packages
                    e.g. gogo import package1 package2 package3
    gogo init       init workspace directory according to package.yaml file
    gogo install    install all import packages according to package.yaml file
    gogo version    print gogo version information
    gogo - <cmd>    run any go command
                    e.g. gogo - env
                         gogo - vet
`)

}
