package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/mgutz/ansi"
)

var phosphorize = ansi.ColorFunc("gray+h")

type Command struct {
	dir    string
	env    []string
	status int
	cmd    *exec.Cmd
}

func NewCommand(path string, args ...string) (*Command, error) {
	var fullPath = ""
	fullPath, err := exec.LookPath(path)
	if err != nil {
		return nil, err
	}
	log.Println(phosphorize("exec:     " + fullPath + " " + strings.Join(args, " ")))
	return &Command{
		cmd: exec.Command(fullPath, args...),
		env: os.Environ(),
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

func (c *Command) RunAndGetOutputs() ([]byte, error) {
	return c.cmd.CombinedOutput()
}

func (c *Command) Run() {
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr
	c.cmd.Stdin = os.Stdin
	if err := c.cmd.Run(); err != nil {
		log.Println(phosphorize("exec error: " + err.Error()))
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
