package sshconn

import (
	"errors"
	"os/exec"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

// proxyCommandProcess runs a command line the way OpenSSH for Windows does:
// the program is its first word, quoted or not, and the program reads the rest
// of the line itself, so a path such as "C:\Program Files\..." keeps its
// backslashes. No console window opens for it.
func proxyCommandProcess(line string) (*exec.Cmd, error) {
	program := firstCommandWord(line)
	if program == "" {
		return nil, errors.New("the proxy command names no program")
	}
	process := exec.Command(program)
	process.SysProcAttr = &syscall.SysProcAttr{CmdLine: line, HideWindow: true, CreationFlags: windows.CREATE_NO_WINDOW}
	return process, nil
}

// firstCommandWord is the program of a command line, as Windows reads it: up to
// the closing quote when it starts with one, else up to the first blank.
func firstCommandWord(line string) string {
	line = strings.TrimLeft(line, " \t")
	if rest, quoted := strings.CutPrefix(line, `"`); quoted {
		program, _, _ := strings.Cut(rest, `"`)
		return program
	}
	if end := strings.IndexAny(line, " \t"); end >= 0 {
		return line[:end]
	}
	return line
}
