package util

import (
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

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
	Pwd               string
	PwdUnderWorkspace string
	Workspace         string
	Vendor            string
}

// AddImport 添加依赖
func (pkg *PackageInfo) AddImport(name string, version string) {
	pkg.Import = append(pkg.Import, &ImportInfo{
		Package: name,
		Version: version,
	})
}

// GetImportInfo 获取指定依赖模块信息
func (pkg *PackageInfo) GetImportInfo(name string) *ImportInfo {
	for _, item := range pkg.Import {
		if item.Package == name {
			return item
		}
	}
	return nil
}

// GetPackageInfo 从指定目录加载package信息
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
		Pwd:               dir,
		PwdUnderWorkspace: filepath.Join(dir, "_workspace", "src", pkg.Package),
		Workspace:         filepath.Join(dir, "_workspace"),
		Vendor:            filepath.Join(dir, "vendor"),
	}
	return pkg, nil
}

// GetPackageInfoFromCurrentDir 从当前目录加载package信息
func GetPackageInfoFromCurrentDir() (*PackageInfo, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return GetPackageInfo(dir)
}

// SavePackageInfo 保存package信息到指定目录
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

// SavePackageInfoToCurrentDir 保存package信息到当前目录
func SavePackageInfoToCurrentDir(pkg *PackageInfo) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	return SavePackageInfo(dir, pkg)
}

// GetAllGoFilesFromDir 查找指定目录下的所有go文件
func GetAllGoFilesFromDir(dir string) []string {
	var list []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".go") {
			list = append(list, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return list
}

// ReadFileContents 读取文件内容
func ReadFileContents(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(data)
}

func stripQuotes(str string) string {
	if string(str[0]) == `"` {
		str = string(str[1:])
	}
	if string(str[len(str)-1]) == `"` {
		str = string(str[0 : len(str)-1])
	}
	return str
}

// isExternalPackage 判断是否为外部依赖
func isExternalPackage(name string) (bool, string) {
	name = stripQuotes(name)
	items := strings.Split(name, "/")
	if len(items) < 1 {
		return false, ""
	}
	if !strings.Contains(items[0], ".") {
		return false, name
	}
	if len(items) <= 3 {
		return true, name
	}
	return true, strings.Join(items[0:3], "/")
}

// GetSourceImportsFromDir 从指定目录解析go源码并获取import的模块
func GetSourceImportsFromDir(dir string) []string {
	var keys []string
	ret := map[string]string{}
	files := GetAllGoFilesFromDir(dir)
	for _, file := range files {
		fset := token.NewFileSet()
		src := ReadFileContents(file)
		f, err := parser.ParseFile(fset, "", src, parser.ImportsOnly)
		if err != nil {
			log.Fatal(err)
			return keys
		}
		for _, s := range f.Imports {
			if ok, name := isExternalPackage(s.Path.Value); ok {
				ret[name] = name
			}
		}
	}
	for key := range ret {
		keys = append(keys, key)
	}
	return keys
}
