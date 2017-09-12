package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/leizongmin/gogo/util"
	"github.com/mgutz/ansi"
)

var phosphorize = ansi.ColorFunc("gray+h")
var debugMode = false

type execFunctionType func(dir string, name string, args ...string)
type execFunctionWithOutputType func(dir string, name string, args ...string) string

// SetDebug 设置调试输出模式
func SetDebug(enable bool) {
	debugMode = enable
}

// 调试输出
func debugPrintln(line string) {
	if debugMode {
		log.Println(phosphorize(line))
	}
}

// 获取包信息及 exec 函数
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

// 获取 exec 函数
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

// 合并字符串数组
func combineStringArray(a []string, b []string) []string {
	ret := make([]string, len(a)+len(b))
	copy(ret, a)
	copy(ret[len(a):], b)
	return ret
}

// 在字符串数组中查找字符串
func findInStringArray(arr []string, str string) int {
	for i, v := range arr {
		if v == str {
			return i
		}
	}
	return -1
}

// 检查文件或目录是否存在
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

// 判断 _workspace 目录是否存在
func isWorkspaceDirExists(dir string) bool {
	ret, err := exists(filepath.Join(dir, "_workspace"))
	if err != nil {
		log.Fatal(err)
	}
	return ret
}
