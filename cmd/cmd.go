package cmd

import (
	"fmt"
	"log"

	"path/filepath"

	"github.com/leizongmin/gogo/util"
)

type execFunctionType func(name string, args ...string)

func getPackageInfoAndExec(isVendor bool) (*util.PackageInfo, execFunctionType) {
	pkg, err := util.GetPackageInfoFromCurrentDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("package: ", pkg.Package)

	exec := func(name string, args ...string) {
		cmd, err := util.NewCommand(name, args...)
		if err != nil {
			log.Fatal(err)
		}
		gopath := pkg.Dir.Workspace
		if isVendor {
			gopath = filepath.Join(pkg.Dir.Workspace, "vendor")
		}
		cmd.SetEnv("GOPATH", gopath)
		cmd.SetDir(pkg.Dir.Pwd)
		fmt.Println("pwd: ", pkg.Dir.Pwd)
		fmt.Println("gopath: ", gopath)
		cmd.Run()
	}

	return pkg, exec
}
