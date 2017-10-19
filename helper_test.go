package machineid

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"testing"
)

func Test_protect(t *testing.T) {
	appID := "ms.azur.appX"
	machineID := "1a1238d601ad430cbea7efb0d1f3d92d"
	hash := protect(appID, machineID)
	if hash == "" {
		t.Error("hash is empty")
	}
	rawHash, err := hex.DecodeString(hash)
	if err != nil {
		t.Error(err)
	}
	eq := checkMAC([]byte(appID), rawHash, []byte(machineID))
	if eq == false {
		t.Error("hashes do not match")
	}
	// modify rawHash --> should not match
	rawHash[0] = 0
	eq = checkMAC([]byte(appID), rawHash, []byte(machineID))
	if eq == true {
		t.Error("hashes do match, but shouldn't")
	}
}

func checkMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

func Test_run(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	wantStdout := "hello"
	wantStderr := ""
	if err := run(stdout, stderr, "echo", "hello"); err != nil {
		t.Error(err)
	}
	gotStdout := strings.TrimRight(stdout.String(), "\r\n")
	if gotStdout != wantStdout {
		t.Errorf("run() = %v, want %v", gotStdout, wantStdout)
	}
	if gotStderr := stderr.String(); gotStderr != wantStderr {
		t.Errorf("run() = %v, want %v", gotStderr, wantStderr)
	}
}

func Test_run_unknown(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	err := run(stdout, stderr, "echolo", "hello")
	if err == nil {
		t.Error("expected error, got none")
	}
	if strings.Contains(err.Error(), "executable file not found") == false {
		t.Error("unexpected error, expected exec not found")
	}
}
