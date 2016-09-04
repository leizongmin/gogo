package cmd

import (
	"fmt"
)

var version = "0.0.0"

func Version(args []string) {

	fmt.Printf("gogo version %s\n", version)

}
