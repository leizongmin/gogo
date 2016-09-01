package cmd

import (
	"fmt"
	"log"

	"path/filepath"

	"github.com/leizongmin/gogo/util"
)

func Init(args []string) {

	pkg, err := util.GetPackageInfoFromCurrentDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pkg)

	exec := func(name string, args ...string) {
		cmd, err := util.NewCommand(name, args...)
		if err != nil {
			log.Fatal(err)
		}
		cmd.SetEnv("GOPATH", pkg.Dir.Workspace)
		cmd.SetDir(pkg.Dir.Pwd)
		cmd.Run()
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

	exec("mkdir", "-p", filepath.Join(workspace, "src", filepath.Dir(pkg.Package)))
	exec("ln", "-s", pwd, filepath.Join(workspace, "src", pkg.Package))

	exec("mkdir", "-p", filepath.Join(pwd, "vendor"))
	exec("mkdir", "-p", filepath.Join(workspace, "vendor"))
	exec("ln", "-s", filepath.Join(pwd, "vendor"), filepath.Join(workspace, "vendor", "src"))

	exec("mkdir", "-p", filepath.Join(workspace, "vendor", "pkg"))
	exec("ln", "-s", filepath.Join(workspace, "vendor", "pkg"), filepath.Join(workspace, "pkg"))

}