package main

import "os"

import "github.com/leizongmin/gogo/cmd"

func main() {

	if len(os.Args) < 2 {
		cmd.Help([]string{})
		return
	}

	var args = os.Args[2:]
	switch os.Args[1] {
	case "build":
		cmd.Build(args)
	case "clean":
		cmd.Clean(args)
	case "-":
		cmd.Go(args)
	case "import":
		cmd.Import(args)
	case "init":
		cmd.Init(args)
	case "install":
		cmd.Install(args)
	case "version":
		cmd.Version(args)
	case "help":
		cmd.Help(args)
	default:
		cmd.Help(args)
	}

}
