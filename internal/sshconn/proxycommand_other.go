//go:build !windows

package sshconn

import "os/exec"

// proxyCommandProcess runs a command line the way OpenSSH does: the shell reads
// it, and exec leaves the program in the shell's place.
func proxyCommandProcess(line string) (*exec.Cmd, error) {
	return exec.Command("/bin/sh", "-c", "exec "+line), nil
}
