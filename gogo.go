package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/leizongmin/gogo/cmd"
	"github.com/leizongmin/gogo/util"
	"github.com/mgutz/ansi"
)

var red = ansi.ColorFunc("red+h")

func main() {

	if len(os.Args) < 2 {
		cmd.Help([]string{})
		return
	}

	debug := strings.ToLower(os.Getenv("GOGO_DEBUG"))
	if debug == "1" || debug == "true" || debug == "yes" {
		cmd.SetDebug(true)
	} else {
		cmd.SetDebug(false)
	}

	var args = os.Args[2:]
	switch os.Args[1] {
	case "build":
		cmd.Build(args)
	case "dev":
		cmd.Dev(args)
	case "clean":
		cmd.Clean(args)
	case "-":
		cmd.Go(args)
	case "run":
		cmd.Run(args)
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
		var execCmd, err = util.LookupExecPath(os.Args[1])
		if err != nil {
			fmt.Println(red(err.Error()))
			fmt.Println()
			cmd.Help(args)
		} else {
			cmd.Run(append([]string{execCmd}, args...))
		}
	}

}
