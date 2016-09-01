package main

import "os"

import "github.com/leizongmin/gogo/cmd"

func main() {

	if len(os.Args) < 2 {
		cmd.Help([]string{})
		return
	}

	var args = os.Args[2:]
	switch os.Args[1] {
	case "init":
		cmd.Init(args)
	case "clean":
		cmd.Clean(args)
	case "vendor":
		cmd.Vendor(args)
	case "build":
		cmd.Build(args)
	case "run":
		cmd.Run(args)
	case "version":
		cmd.Version(args)
	case "help":
		cmd.Help(args)
	default:
		cmd.Help(args)
	}

}

//func test() {
//	a, err := util.NewCommand("node", "-p", "console.log(process.env)")
//	//	a, err := util.NewCommand("tree", "-d")
//	if err != nil {
//		log.Fatal(err)
//	}
//	dir, err := os.Getwd()
//	if err != nil {
//		log.Fatal(err)
//	}
//	a.SetEnv("GOPATH", dir)
//	a.Run()
//}
