package cmd

import (
	"fmt"
	"path/filepath"
)

// Build 构建项目
func Build(args []string) {
	pkg, exec := getPackageInfoAndExec(false)
	if !isWorkspaceDirExists(pkg.Dir.Pwd) {
		fmt.Println(`
"_workspace" directory doesn't exists, please run "gogo init" before.
		`)
		return
	}
	out := fmt.Sprintf("bin/%s", filepath.Base(pkg.Package))
	exec.setDir(pkg.Dir.Pwd)
	exec.run("go", "build", "-o", out, pkg.Package)
	fmt.Println("OK")
}

// BuildHelp 命令帮助
func BuildHelp(args []string) {
	fmt.Println(`
usage: gogo build

compile the current project.
	`)
}
