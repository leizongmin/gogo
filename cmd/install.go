package cmd

import "fmt"
import "github.com/leizongmin/gogo/util"
import "path/filepath"

func Install(args []string) {

	pkg, exec := getPackageInfoAndExec(true)

	if !isWorkspaceDirExists(pkg.Dir.Pwd) {
		fmt.Println(`
"_workspace" directory doesn't exists, please run "gogo init" before.
		`)
		return
	}

	if len(pkg.Package) > 0 {
		for _, p := range pkg.Import {
			downloadPackage(pkg, exec, p)
		}
	}

	fmt.Println("\nOK")

}

func InstallHelp(args []string) {
	fmt.Println(`
usage: gogo install

install all import packages according to package.yaml file
	`)
}

func downloadPackage(pkg *util.PackageInfo, exec execFunctionType, info *util.ImportInfo) {
	pkgPath := filepath.Join(pkg.Dir.Pwd, "vendor", info.Package)
	exec(pkg.Dir.Pwd, "rm", "-rf", pkgPath)
	exec(pkg.Dir.Pwd, "git", "clone", "https://"+info.Package+".git", pkgPath)
	if info.Package != "*" && info.Package != "" {
		exec(pkgPath, "git", "checkout", info.Version)
	}
	exec(pkg.Dir.Pwd, "rm", "-rf", filepath.Join(pkgPath, ".git"))
}
