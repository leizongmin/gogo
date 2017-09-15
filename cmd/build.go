package cmd

import (
	"fmt"
	"log"
	"path/filepath"
)

// Build 构建项目
func Build(args []string) {
	pkg, exec := getPackageInfoAndExec(false)
	if !isWorkspaceDirExists(pkg.Dir.Pwd) {
		log.Println(`"_workspace" directory doesn't exists, please run "gogo init" before.`)
		return
	}
	out := fmt.Sprintf("bin/%s", filepath.Base(pkg.Package))
	exec.setDir(pkg.Dir.Pwd)

	buildArgs := append([]string{"build"}, args...)
	buildArgs = append(buildArgs, "-o", out, pkg.Package)
	ret := exec.run("go", buildArgs...)
	fmt.Print(ret)
	log.Println("OK")
}

// BuildHelp 命令帮助
func BuildHelp(args []string) {
	fmt.Println(`
Usage: gogo build

compile the current project.
	`)
}
