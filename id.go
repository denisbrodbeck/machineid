// Package machineid provides support for reading the unique machine id of most host OS's (without admin privileges).
package machineid // import "github.com/denisbrodbeck/machineid"

import (
	"io"
	"os"
	"os/exec"
)

// ID returns the platform specific machine id of the current host.
func ID() (string, error) {
	return machineID()
}

// run wraps `exec.Command` with easy access to stdout and stderr.
func run(stdout, stderr io.Writer, cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Stdin = os.Stdin
	c.Stdout = stdout
	c.Stderr = stderr
	return c.Run()
}
