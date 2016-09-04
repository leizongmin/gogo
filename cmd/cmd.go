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

func combineStringArray(a []string, b []string) []string {
	ret := make([]string, len(a)+len(b))
	copy(ret, a)
	copy(ret[len(a):], b)
	return ret
}

func findInStringArray(arr []string, str string) int {
	for i, v := range arr {
		if v == str {
			return i
		}
	}
	return -1
}
