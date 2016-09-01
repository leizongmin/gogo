package cmd

import "path/filepath"

func Clean(args []string) {

	pkg, exec := getPackageInfoAndExec(false)

	// `rm -rf ${workspace}`,
	// `rm -rf ${pwd}/vendor`,
	// `rm -rf ${pwd}/bin`,

	workspace := pkg.Dir.Workspace
	pwd := pkg.Dir.Pwd

	exec("rm", "-rf", workspace)
	exec("rm", "-rf", filepath.Join(pwd, "vendor"))
	exec("rm", "-rf", filepath.Join(pwd, "bin"))

}
