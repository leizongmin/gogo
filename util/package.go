package util

import (
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type PackageInfo struct {
	Package   string       `yaml:"package"`
	Version   string       `yaml:"version"`
	Author    string       `yaml:"author"`
	Homepage  string       `yaml:"homepage"`
	Import    []ImportInfo `yaml:"import"`
	Workspace WorkspaceInfo
}

type ImportInfo struct {
	Package string `yaml:"package"`
	Version string `yaml:"version"`
}

type WorkspaceInfo struct {
	ProjectDir          string
	VendorDir           string
	VirtualWorkspaceDir string
	VirtualVendorDir    string
	VirtualPkgDir       string
	BinDir              string
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
	pkg.Workspace = WorkspaceInfo{
		ProjectDir:          dir,
		VendorDir:           filepath.Join(dir, "vendor"),
		VirtualWorkspaceDir: filepath.Join(dir, "_workspace"),
		VirtualPkgDir:       filepath.Join(dir, "_workspace", "pkg"),
		VirtualVendorDir:    filepath.Join(dir, "_workspace", "vendor"),
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
