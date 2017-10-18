// +build freebsd netbsd openbsd

package machineid

import (
	"bytes"
	"os"
	"strings"
)

// machineID returns the uuid returned by `kenv -q smbios.system.uuid`.
// If there is an error running the commad an empty string is returned.
func machineID() (string, error) {
	buf := &bytes.Buffer{}
	err := run(buf, os.Stderr, "kenv", "-q", "smbios.system.uuid")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(buf.String()), nil
}
