package cmd

import (
	"fmt"
	"log"

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
		cmd.SetEnv("GOPATH", pkg.Workspace.VirtualWorkspaceDir)
		cmd.SetDir(pkg.Workspace.ProjectDir)
		cmd.Run()
	}

	exec("mkdir", "-p", pkg.Workspace.VirtualWorkspaceDir)

}
