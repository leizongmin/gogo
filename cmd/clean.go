package cmd

import (
	"fmt"
	"path/filepath"
)

func Clean(args []string) {

	pkg, exec := getPackageInfoAndExec(false)

	// `rm -rf ${workspace}`,
	// `rm -rf ${pwd}/vendor`,
	// `rm -rf ${pwd}/bin`,

	workspace := pkg.Dir.Workspace
	pwd := pkg.Dir.Pwd

	exec(pwd, "rm", "-rf", workspace)
	exec(pwd, "rm", "-rf", filepath.Join(pwd, "vendor"))
	exec(pwd, "rm", "-rf", filepath.Join(pwd, "bin"))

	fmt.Println("OK")

}

func CleanHelp(args []string) {
	fmt.Println(`
usage: gogo clean

remove object files from current project directory.
	`)
}
