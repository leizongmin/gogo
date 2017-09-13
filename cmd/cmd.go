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

type execObject struct {
	dir string
	env map[string]string
}

func createExec() *execObject {
	return &execObject{
		dir: ".",
		env: make(map[string]string),
	}
}

func (e *execObject) setDir(dir string) {
	e.dir = dir
}

func (e *execObject) setEnv(name string, value string) {
	e.env[name] = value
}

func (e *execObject) getCmd(bin string, args ...string) *util.Command {
	cmd, err := util.NewCommand(bin, args...)
	if err != nil {
		log.Fatal(err)
	}
	cmd.SetDebugPrintln(debugPrintln)
	cmd.SetDir(e.dir)
	for k, v := range e.env {
		cmd.SetEnv(k, v)
	}
	return cmd
}

func (e *execObject) run(bin string, args ...string) string {
	cmd := e.getCmd(bin, args...)
	ret, err := cmd.RunAndGetOutputs()
	if err != nil {
		debugPrintln(err.Error())
	} else {
		debugPrintln("Success")
	}
	return string(ret)
}

// 获取包信息及 exec 函数
func getPackageInfoAndExec(isVendor bool) (*util.PackageInfo, *execObject) {
	pkg, err := util.GetPackageInfoFromCurrentDir()
	if err != nil {
		log.Fatal(err)
	}
	debugPrintln("package:  " + pkg.Package)

	gopath := pkg.Dir.Workspace
	if isVendor {
		gopath = filepath.Join(pkg.Dir.Workspace, "vendor")
	}
	exec := createExec()
	exec.setEnv("GOPATH", gopath)

	return pkg, exec
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
