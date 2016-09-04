package cmd

func Go(args []string) {

	_, exec := getPackageInfoAndExec(false)

	exec("go", args...)

}
