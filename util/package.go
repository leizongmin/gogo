package util

import (
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type PackageInfo struct {
	Package  string        `yaml:"package"`
	Version  string        `yaml:"version"`
	Author   string        `yaml:"author"`
	Homepage string        `yaml:"homepage"`
	Import   []*ImportInfo `yaml:"import"`
	Dir      *DirInfo      `yaml:"dir,omitempty"`
}

type ImportInfo struct {
	Package string `yaml:"package"`
	Version string `yaml:"version"`
}

type DirInfo struct {
	Pwd       string
	Workspace string
}

func (pkg *PackageInfo) AddImport(name string, version string) {
	pkg.Import = append(pkg.Import, &ImportInfo{
		Package: name,
		Version: version,
	})
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
	pkg.Dir = &DirInfo{
		Pwd:       dir,
		Workspace: filepath.Join(dir, "_workspace"),
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

func SavePackageInfo(dir string, pkg *PackageInfo) error {
	file := filepath.Join(dir, "package.yaml")
	dirInfo := pkg.Dir
	pkg.Dir = nil
	fileData, err := yaml.Marshal(pkg)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(file, fileData, 0664); err != nil {
		return err
	}
	pkg.Dir = dirInfo
	return nil
}

func SavePackageInfoToCurrentDir(pkg *PackageInfo) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	return SavePackageInfo(dir, pkg)
}
