package cmd

import "path/filepath"
import "fmt"

func Install(args []string) {

	pkg, exec := getPackageInfoAndExec(false)

	// `go build -o ${out} ${pkg}`

	out := fmt.Sprintf("bin/%s", filepath.Base(pkg.Package))

	exec("go", "build", "-o", out, pkg.Package)

	fmt.Println("\nOK")

}

func InstallHelp(args []string) {
	fmt.Println(`
usage: gogo install

install all import packages according to package.yaml file
	`)
}
