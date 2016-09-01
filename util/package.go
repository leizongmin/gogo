package util

import (
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type PackageInfo struct {
	Package  string       `yaml:"package"`
	Version  string       `yaml:"version"`
	Author   string       `yaml:"author"`
	Homepage string       `yaml:"homepage"`
	Import   []ImportInfo `yaml:"import"`
	Dir      DirInfo
}

type ImportInfo struct {
	Package string `yaml:"package"`
	Version string `yaml:"version"`
}

type DirInfo struct {
	Workspace string
	Pwd       string
}

func GetPackageInfo(dir string) (*PackageInfo, error) {
	file := filepath.Join(dir, "package.yaml")
	fileData, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	pkg := &PackageInfo{}
	err = yaml.Unmarshal(fileData, &pkg)
	if err != nil {
		return nil, err
	}
	pkg.Dir = DirInfo{
		Workspace: filepath.Join(dir, "_workspace"),
		Pwd:       dir,
	}
	return pkg, nil
}

func GetPackageInfoFromCurrentDir() (*PackageInfo, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return GetPackageInfo(dir)
}
