package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mgutz/ansi"
)

var phosphorize = ansi.ColorFunc("gray+h")

type Command struct {
	execPath     string
	args         []string
	dir          string
	env          []string
	status       int
	cmd          *exec.Cmd
	debugPrintln func(string)
}

func NewCommand(path string, args ...string) (*Command, error) {
	var fullPath = ""
	fullPath, err := exec.LookPath(path)
	if err != nil {
		return nil, err
	}
	return &Command{
		execPath: fullPath,
		args:     args,
		cmd:      exec.Command(fullPath, args...),
		env:      os.Environ(),
	}, nil
}

func (c *Command) SetEnv(name string, value string) {
	c.env = append(filterEnv(c.env, name), fmt.Sprintf("%s=%s", name, value))
	c.cmd.Env = c.env
}

func (c *Command) SetEnvLine(line string) {
	c.env = append(c.env, line)
	c.cmd.Env = c.env
}

func (c *Command) SetDir(dir string) {
	c.dir = dir
	c.cmd.Dir = c.dir
}

func (c *Command) SetDebugPrintln(fn func(string)) {
	c.debugPrintln = fn
}

func (c *Command) RunAndGetOutputs() ([]byte, error) {
	if c.debugPrintln != nil {
		c.debugPrintln("exec:     " + c.execPath + " " + strings.Join(c.args, " "))
	}
	return c.cmd.CombinedOutput()
}

func (c *Command) Run() {
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr
	c.cmd.Stdin = os.Stdin
	if c.debugPrintln != nil {
		c.debugPrintln("exec:     " + c.execPath + " " + strings.Join(c.args, " "))
	}
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
