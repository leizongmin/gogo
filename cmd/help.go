package cmd

import (
	"fmt"
)

func Help(args []string) {

	Version(args)

	if len(args) < 1 {
		printHelpIndex()
		return
	}

	switch args[0] {
	case "build":
		BuildHelp(args[1:])
	case "clean":
		CleanHelp(args[1:])
	case "help":
		fmt.Println("print the help information\n")
	case "import":
		ImportHelp(args[1:])
	case "init":
		InitHelp(args[1:])
	case "install":
		InstallHelp(args[1:])
	case "version":
		VersionHelp(args[1:])
	case "-":
		GoHelp(args[1:])
	default:
		fmt.Printf("Unknown help topic \"%s\".  Run \"gogo help\".\n", args[0])
	}

}

func printHelpIndex() {
	fmt.Printf(`
Usage:

    gogo command [arguments]

The commands are:

    gogo build      compile the current project
    gogo clean      remove vendor and workspace directory
    gogo help       print the help information
    gogo import     add import packages and save to package.yaml file
                    e.g. gogo import package1 package2 package3
    gogo init       init workspace directory according to package.yaml file
    gogo install    install all import packages according to package.yaml file
    gogo version    print gogo version information
    gogo - <cmd>    run any go command
                    e.g. gogo - env
                         gogo - vet

Use "gogo help [command]" for more information about a command.
	`)
}
