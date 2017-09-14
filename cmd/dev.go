package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/leizongmin/gogo/util"
)

// Dev 构建项目并运行
func Dev(args []string) {
	pkg, exec := getPackageInfoAndExec(false)
	if !isWorkspaceDirExists(pkg.Dir.Pwd) {
		log.Println(`"_workspace" directory doesn't exists, please run "gogo init" before.`)
		return
	}
	out := fmt.Sprintf("bin/%s", filepath.Base(pkg.Package))
	exec.setDir(pkg.Dir.Pwd)
	ret := exec.run("go", "build", "-o", out, pkg.Package)
	fmt.Print(ret)
	log.Println("OK")

	cmd, err := util.NewCommand(out, args...)
	if err != nil {
		log.Fatal(err)
	}
	cmd.SetDir(pkg.Dir.Pwd)
	log.Printf("Run %s %s", out, strings.Join(args, " "))
	cmd.Run()
}

// DevHelp 命令帮助
func DevHelp(args []string) {
	fmt.Println(`
Usage: gogo dev

compile the current project and run.
	`)
}
