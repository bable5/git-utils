package git

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
	"syscall"
)

type UserStoryNotFound struct {
	msg string
}

func NewUserStoryNotFound(msg string) error {
	return &UserStoryNotFound{
		msg: msg,
	}
}

func (m *UserStoryNotFound) Error() string {
	return m.msg
}

func AddToConfig(configEntry, value string) error {
	var stderr bytes.Buffer

	args := []string{"config", "--replace-all", configEntry, value}
	err := RunGitCommand(ioutil.Discard, &stderr, args)

	if err != nil {
		fmt.Println(stderr.String())
	}

	return err
}

func RemoveFromConfig(configEntry string) error {
	args := []string{"config", "--unset", configEntry}
	return RunGitCommand(ioutil.Discard, ioutil.Discard, args)
}

func GetFromConfig(configEntry string) (string, error) {
	var stdout bytes.Buffer

	err := RunGitCommand(&stdout, ioutil.Discard, []string{"config", "--null", "--get", configEntry})
	if exitError, ok := err.(*exec.ExitError); ok {
		if waitStatus, ok := exitError.Sys().(syscall.WaitStatus); ok {
			if waitStatus.ExitStatus() == 1 {
				return "", NewUserStoryNotFound(fmt.Sprintf("%s not found", configEntry))
			}
		}
		return "", err
	}

	return strings.TrimRight(stdout.String(), "\000"), nil
}

func RunGitCommand(stdout, stderr io.Writer, args []string) error {
	cmd := MakeGitCommand(args)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	return cmd.Run()
}

func MakeGitCommand(args []string) *exec.Cmd {
	return exec.Command("git", args...)
}
