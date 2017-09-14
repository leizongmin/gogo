package cmd

import (
	"fmt"
	"log"
	"path/filepath"
)

// Clean 清理项目
func Clean(args []string) {
	pkg, exec := getPackageInfoAndExec(false)
	workspace := pkg.Dir.Workspace
	pwd := pkg.Dir.Pwd
	exec.setDir(pwd)
	exec.run("rm", "-rf", workspace)
	exec.run("rm", "-rf", filepath.Join(pwd, "vendor"))
	exec.run("rm", "-rf", filepath.Join(pwd, "bin"))
	log.Println("OK")
}

// CleanHelp 命令帮助
func CleanHelp(args []string) {
	fmt.Println(`
Usage: gogo clean

remove object files from current project directory.
	`)
}
