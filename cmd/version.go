package cmd

import (
	"fmt"
)

var version = "0.0.1"

func Version(args []string) {

	fmt.Printf("gogo version %s\n", version)

}

func VersionHelp(args []string) {
	fmt.Println(`
usage: gogo version

print gogo version information.
	`)
}
