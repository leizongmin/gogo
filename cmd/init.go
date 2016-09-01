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

}
