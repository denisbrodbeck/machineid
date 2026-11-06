// +build aix

package machineid

import (
	"bytes"
	"os"
)

func machineID() (string, error) {
	buf := &bytes.Bytes{}
	err := run(buf, os.Stderr, "/usr/sbin/lsattr", "-E", "-l", "sys0", "-a", "os_uuid", "-F", "value")
	if err != nil {
		return "", err
	}
	return trim(buf.String()), nil
}
