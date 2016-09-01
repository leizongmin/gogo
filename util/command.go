package util

import (
	"fmt"
	"os"
	"os/exec"
)

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
	fmt.Println("exec:", fullPath, args)
	return &Command{
		cmd: exec.Command(fullPath, args...),
		env: os.Environ(),
	}, nil
}

func (c *Command) SetEnv(name string, value string) {
	c.env = append(c.env, fmt.Sprintf("%s=%s", name, value))
}

func (c *Command) SetDir(dir string) {
	c.dir = dir
}

func (c *Command) Run() {
	c.cmd.Env = c.env
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr
	c.cmd.Stdin = os.Stdin
	c.cmd.Dir = c.dir
	if err := c.cmd.Run(); err != nil {
		fmt.Println("exec:", err)
	}
}
