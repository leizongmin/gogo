package cmd

import (
	"fmt"
	"log"

	"path/filepath"

	"os"

	"github.com/leizongmin/gogo/util"
	"github.com/mgutz/ansi"
)

var phosphorize = ansi.ColorFunc("gray+h")

type execFunctionType func(dir string, name string, args ...string)
type execFunctionWithOutputType func(dir string, name string, args ...string) string

func getPackageInfoAndExec(isVendor bool) (*util.PackageInfo, execFunctionType) {
	pkg, err := util.GetPackageInfoFromCurrentDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(phosphorize("package:  " + pkg.Package))

	exec := func(dir string, name string, args ...string) {
		cmd, err := util.NewCommand(name, args...)
		if err != nil {
			log.Fatal(err)
		}
		gopath := pkg.Dir.Workspace
		if isVendor {
			gopath = filepath.Join(pkg.Dir.Workspace, "vendor")
		}
		cmd.SetEnv("GOPATH", gopath)
		cmd.SetDir(dir)
		fmt.Println(phosphorize("PWD:      " + dir))
		fmt.Println(phosphorize("GOPATH:   " + gopath))
		cmd.Run()
		fmt.Println(phosphorize("Success"))
	}

	return pkg, exec
}

func getExec() execFunctionWithOutputType {
	return func(dir string, name string, args ...string) string {
		cmd, err := util.NewCommand(name, args...)
		if err != nil {
			log.Fatal(err)
		}
		cmd.SetDir(dir)
		fmt.Println(phosphorize("PWD:      " + dir))
		ret, err := cmd.RunAndGetOutputs()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(phosphorize("Success"))
		return string(ret)
	}
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

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func isWorkspaceDirExists(dir string) bool {
	ret, err := exists(filepath.Join(dir, "_workspace"))
	if err != nil {
		log.Fatal(err)
	}
	return ret
}
