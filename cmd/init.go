package cmd

import (
	"fmt"
	"log"
	"path/filepath"
)

func Init(args []string) {

	pkg, exec := getPackageInfoAndExec(false)

	if isWorkspaceDirExists(pkg.Dir.Pwd) {
		log.Println(`"_workspace" directory already exists, please run "gogo clean" before.`)
		return
	}

	// `mkdir -p ${workspace}/src/${pkgParent}`,
	// `ln -s ${pwd} ${workspace}/src/${pkg}`,
	// `mkdir -p ${pwd}/vendor`,
	// `mkdir -p ${workspace}/vendor`,
	// `ln -s ${pwd}/vendor ${workspace}/vendor/src`,
	// `mkdir -p ${workspace}/vendor/pkg`,
	// `ln -s ${workspace}/vendor/pkg ${workspace}/pkg`,

	workspace := pkg.Dir.Workspace
	pwd := pkg.Dir.Pwd

	exec(pwd, "mkdir", "-p", filepath.Join(workspace, "src", filepath.Dir(pkg.Package)))
	exec(pwd, "ln", "-s", pwd, filepath.Join(workspace, "src", pkg.Package))

	exec(pwd, "mkdir", "-p", filepath.Join(pwd, "vendor"))
	exec(pwd, "mkdir", "-p", filepath.Join(workspace, "vendor"))
	exec(pwd, "ln", "-s", filepath.Join(pwd, "vendor"), filepath.Join(workspace, "vendor", "src"))

	exec(pwd, "mkdir", "-p", filepath.Join(workspace, "vendor", "pkg"))
	exec(pwd, "ln", "-s", filepath.Join(workspace, "vendor", "pkg"), filepath.Join(workspace, "pkg"))

	log.Println("OK")

}

func InitHelp(args []string) {
	fmt.Println(`
usage: gogo init

init workspace directory according to package.yaml file.
	`)
}
