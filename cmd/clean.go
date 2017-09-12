package cmd

import (
	"fmt"
	"path/filepath"
)

// Clean 清理项目
func Clean(args []string) {
	pkg, exec := getPackageInfoAndExec(false)
	workspace := pkg.Dir.Workspace
	pwd := pkg.Dir.Pwd

	exec(pwd, "rm", "-rf", workspace)
	exec(pwd, "rm", "-rf", filepath.Join(pwd, "vendor"))
	exec(pwd, "rm", "-rf", filepath.Join(pwd, "bin"))
	fmt.Println("OK")
}

// CleanHelp 命令帮助
func CleanHelp(args []string) {
	fmt.Println(`
usage: gogo clean

remove object files from current project directory.
	`)
}
