package cmd

import (
	"log"
	"path/filepath"
	"regexp"

	"github.com/leizongmin/gogo/util"
)

// Install 安装依赖模块
func Install(args []string) {
	pkg, exec := getPackageInfoAndExec(true)
	if !isWorkspaceDirExists(pkg.Dir.Pwd) {
		log.Println(`
"_workspace" directory doesn't exists, please run "gogo init" before.
		`)
		return
	}
	if len(pkg.Package) > 0 {
		for _, p := range pkg.Import {
			downloadPackage(pkg, exec, p)
		}
	}
	log.Println("OK")
}

// InstallHelp 命令帮助
func InstallHelp(args []string) {
	log.Println(`
usage: gogo install

install all import packages according to package.yaml file
	`)
}

// 下载模块
func downloadPackage(pkg *util.PackageInfo, exec *execObject, info *util.ImportInfo) {
	pkgPath := filepath.Join(pkg.Dir.Pwd, "vendor", info.Package)
	log.Printf("Downloading package %s\n", info.Package)
	if isGitRepository(pkgPath) {
		exec.setDir(pkgPath)
		exec.run("git", "reset", "--hard", "HEAD")
		exec.run("git", "pull")
	} else {
		exec.setDir(pkg.Dir.Pwd)
		exec.run("rm", "-rf", pkgPath)
		exec.run("git", "clone", "https://"+info.Package+".git", pkgPath)
	}
	if info.Package != "*" && info.Package != "" {
		exec.setDir(pkgPath)
		exec.run("git", "checkout", info.Version)
	}
	info.Version = getLastGitCommitHash(pkgPath)
	log.Printf("Package %s at %s\n", info.Package, info.Version)
}

// 判断目录是否为 Git 仓库
func isGitRepository(dir string) bool {
	if !checkPathExists(dir) {
		return false
	}
	if !checkPathExists(filepath.Join(dir, ".git")) {
		return false
	}
	return true
}

// 获取当前 Git 仓库最后一次提交的 hash
func getLastGitCommitHash(dir string) string {
	if !isGitRepository(dir) {
		return ""
	}
	exec := createExec()
	exec.setDir(dir)
	stdout := exec.run("git", "log", "-n", "1")
	debugPrintln(stdout)

	reg := regexp.MustCompile(`[a-z0-9]{40}`)
	ret := reg.FindAllString(stdout, -1)
	if len(ret) > 0 {
		return ret[0]
	}
	return ""
}

// 检查文件路径是否存在
func checkPathExists(path string) bool {
	ret, err := exists(path)
	if err != nil {
		log.Println(err)
	}
	return ret
}
