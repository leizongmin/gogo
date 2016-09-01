package cmd

import (
	"fmt"
	"log"

	"path/filepath"

	"github.com/leizongmin/gogo/util"
)

func Clean(args []string) {

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

	// `rm -rf ${workspace}`,
	// `rm -rf ${pwd}/vendor`,
	// `rm -rf ${pwd}/bin`,

	workspace := pkg.Dir.Workspace
	pwd := pkg.Dir.Pwd

	exec("rm", "-rf", workspace)
	exec("rm", "-rf", filepath.Join(pwd, "vendor"))
	exec("rm", "-rf", filepath.Join(pwd, "bin"))

}
