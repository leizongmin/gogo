package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/leizongmin/gogo/util"
)

// Import 添加模块依赖
func Import(args []string) {
	pkg, exec := getPackageInfoAndExec(true)
	if !isWorkspaceDirExists(pkg.Dir.Pwd) {
		log.Println(`
"_workspace" directory doesn't exists, please run "gogo init" before.
		`)
		return
	}

	count := 0
	packages := make([]string, len(pkg.Import))
	for i, p := range pkg.Import {
		packages[i] = p.Package
	}
	for _, name := range args {
		if strings.Index(name, "-") != 0 && findInStringArray(packages, name) == -1 {
			pkg.AddImport(name, "*")
			count++
		}
	}
	for _, p := range pkg.Import {
		downloadPackage(pkg, exec, p)
	}
	for {
		log.Println("Find implicit dependences...")
		c := downloadAllImplicitDependences(pkg, exec)
		count += c
		if err := util.SavePackageInfoToCurrentDir(pkg); err != nil {
			log.Fatal(err)
		}
		if c < 1 {
			break
		}
	}

	if count > 0 {
		fmt.Printf("Save %d new import(s) to package.yaml file\n", count)
	}
	if err := util.SavePackageInfoToCurrentDir(pkg); err != nil {
		log.Fatal(err)
	}

	log.Println("OK")
}

// ImportHelp 命令帮助
func ImportHelp(args []string) {
	fmt.Println(`
Usage: gogo import package1 package2 package3

add import packages and save to package.yaml file
	`)
}
