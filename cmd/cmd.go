package cmd

import (
	"log"

	"path/filepath"

	"os"

	"github.com/leizongmin/gogo/util"
	"github.com/mgutz/ansi"
)

var phosphorize = ansi.ColorFunc("gray+h")
var debugMode = false

type execFunctionType func(dir string, name string, args ...string)
type execFunctionWithOutputType func(dir string, name string, args ...string) string

func SetDebug(enable bool) {
	debugMode = enable
}

func debugPrintln(line string) {
	if debugMode {
		log.Println(phosphorize(line))
	}
}

func getPackageInfoAndExec(isVendor bool) (*util.PackageInfo, execFunctionType) {
	pkg, err := util.GetPackageInfoFromCurrentDir()
	if err != nil {
		log.Fatal(err)
	}
	debugPrintln("package:  " + pkg.Package)

	exec := func(dir string, name string, args ...string) {
		cmd, err := util.NewCommand(name, args...)
		if err != nil {
			log.Fatal(err)
		}
		gopath := pkg.Dir.Workspace
		if isVendor {
			gopath = filepath.Join(pkg.Dir.Workspace, "vendor")
		}
		cmd.SetDebugPrintln(debugPrintln)
		cmd.SetEnv("GOPATH", gopath)
		cmd.SetDir(dir)
		debugPrintln("PWD:      " + dir)
		debugPrintln("GOPATH:   " + gopath)
		if _, err := cmd.RunAndGetOutputs(); err != nil {
			debugPrintln(err.Error())
		} else {
			debugPrintln("Success")
		}
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
		debugPrintln("PWD:      " + dir)
		ret, err := cmd.RunAndGetOutputs()
		if err != nil {
			log.Fatal(err)
		}
		debugPrintln("Success")
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
