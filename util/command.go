package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mgutz/ansi"
)

var phosphorize = ansi.ColorFunc("gray+h")

// Command 命令结构体
type Command struct {
	execPath     string
	args         []string
	dir          string
	env          []string
	status       int
	cmd          *exec.Cmd
	debugPrintln func(string)
}

// NewCommand 创建新命令
func NewCommand(path string, args ...string) (*Command, error) {
	var execPath = ""
	execPath, err := exec.LookPath(path)
	if err != nil {
		return nil, err
	}
	return &Command{
		execPath:     execPath,
		args:         args,
		cmd:          exec.Command(execPath, args...),
		env:          os.Environ(),
		debugPrintln: nil,
	}, nil
}

// SetEnv 设置环境变量
func (c *Command) SetEnv(name string, value string) {
	c.env = append(filterEnv(c.env, name), fmt.Sprintf("%s=%s", name, value))
	c.cmd.Env = c.env
}

// SetEnvLine 设置环境变量
func (c *Command) SetEnvLine(line string) {
	c.env = append(c.env, line)
	c.cmd.Env = c.env
}

// SetDir 设置工作目录
func (c *Command) SetDir(dir string) {
	c.dir = dir
	c.cmd.Dir = c.dir
}

// SetDebugPrintln 设置调试输出函数
func (c *Command) SetDebugPrintln(fn func(string)) {
	c.debugPrintln = fn
}

func (c *Command) printInfo() {
	if c.debugPrintln != nil {
		c.debugPrintln("pwd:      " + c.dir)
		c.debugPrintln("env:\n" + strings.Join(c.env, "\n"))
		c.debugPrintln("exec:     " + c.execPath + " " + strings.Join(c.args, " "))
	}
}

// RunAndGetOutputs 运行命令，并返回执行完毕后输出的内容
func (c *Command) RunAndGetOutputs() ([]byte, error) {
	c.printInfo()
	return c.cmd.CombinedOutput()
}

// Run 运行命令
func (c *Command) Run() {
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr
	c.cmd.Stdin = os.Stdin
	c.printInfo()
	if err := c.cmd.Run(); err != nil {
		if c.debugPrintln != nil {
			c.debugPrintln("exec error: " + err.Error())
		}
	}
}

func filterEnv(env []string, removeName string) []string {
	removeName = removeName + "="
	ret := []string{}
	for _, v := range env {
		if strings.Index(v, removeName) != 0 {
			ret = append(ret, v)
		}
	}
	return ret
}
