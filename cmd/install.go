package cmd

import (
	"log"
	"path/filepath"
	"regexp"

	"github.com/leizongmin/gogo/util"
)

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

func InstallHelp(args []string) {
	log.Println(`
usage: gogo install

install all import packages according to package.yaml file
	`)
}

func downloadPackage(pkg *util.PackageInfo, exec execFunctionType, info *util.ImportInfo) {
	pkgPath := filepath.Join(pkg.Dir.Pwd, "vendor", info.Package)
	log.Printf("Downloading package %s\n", info.Package)
	if isGitRepository(pkgPath) {
		exec(pkgPath, "git", "reset", "--hard", "HEAD")
		exec(pkgPath, "git", "pull")
	} else {
		exec(pkg.Dir.Pwd, "rm", "-rf", pkgPath)
		exec(pkg.Dir.Pwd, "git", "clone", "https://"+info.Package+".git", pkgPath)
	}
	if info.Package != "*" && info.Package != "" {
		exec(pkgPath, "git", "checkout", info.Version)
	}
	info.Version = getLastGitCommit(pkgPath)
	log.Printf("Package %s at %s\n", info.Package, info.Version)
}

func isGitRepository(dir string) bool {
	if !checkPathExists(dir) {
		return false
	}
	if !checkPathExists(filepath.Join(dir, ".git")) {
		return false
	}
	return true
}

func getLastGitCommit(dir string) string {
	if !isGitRepository(dir) {
		return ""
	}
	exec := getExec()
	stdout := exec(dir, "git", "log", "-n", "1")
	debugPrintln(stdout)

	reg := regexp.MustCompile(`[a-z0-9]{40}`)
	ret := reg.FindAllString(stdout, -1)
	if len(ret) > 0 {
		return ret[0]
	}
	return ""
}

func checkPathExists(path string) bool {
	ret, err := exists(path)
	if err != nil {
		log.Println(err)
	}
	return ret
}
