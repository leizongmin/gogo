package cmd

import "path/filepath"
import "fmt"

func Build(args []string) {

	pkg, exec := getPackageInfoAndExec(false)

	if !isWorkspaceDirExists(pkg.Dir.Pwd) {
		fmt.Println(`
"_workspace" directory doesn't exists, please run "gogo init" before.
		`)
		return
	}

	// `go build -o ${out} ${pkg}`

	out := fmt.Sprintf("bin/%s", filepath.Base(pkg.Package))

	exec(pkg.Dir.Pwd, "go", "build", "-o", out, pkg.Package)

	fmt.Println("\nOK")

}

func BuildHelp(args []string) {
	fmt.Println(`
usage: gogo build

compile the current project.
	`)
}
