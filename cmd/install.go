package cmd

import "fmt"

func Install(args []string) {

	pkg, exec := getPackageInfoAndExec(true)

	if !isWorkspaceDirExists(pkg.Dir.Pwd) {
		fmt.Println(`
"_workspace" directory doesn't exists, please run "gogo init" before.
		`)
		return
	}

	if len(pkg.Package) > 0 {
		packages := make([]string, len(pkg.Import))
		for i, v := range pkg.Import {
			packages[i] = v.Package
		}
		newArgs := combineStringArray([]string{"get"}, packages)
		exec("go", newArgs...)
	}

	fmt.Println("\nOK")

}

func InstallHelp(args []string) {
	fmt.Println(`
usage: gogo install

install all import packages according to package.yaml file
	`)
}
