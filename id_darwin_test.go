// +build darwin

package machineid

import (
	"strings"
	"testing"
)

const sampleOutput = `+-o MacBookPro12,1  <class IOPlatformExpertDevice, id 0x100000112, registered, matched, active, busy 0 (580075 ms), retain 42>
{
  "IOPlatformSystemSleepPolicy" = <534c505402001300841e120004000000001400000004000006000000000000000f2500000000000000004000000040000000100000001000070000$
  "compatible" = <"MacBookPro12,1">
  "version" = <"1.0">
  "board-id" = <"Mac-EEECCCDDD8888AAA">
  "IOInterruptSpecifiers" = (<0900000005000000>)
  "platform-feature" = <0200000000000000>
  "serial-number" = <1111111100000000000000000022222222222d3333333000000000000000000000000000000000000>
  "IOInterruptControllers" = ("io-apic-0")
  "IOPlatformUUID" = "A3344D1DD-1234-22A1-B123-11AB1C11D111"
  "target-type" = <"Mac">
  "clock-frequency" = <00e1f505>
  "manufacturer" = <"Apple Inc.">
  "IOPolledInterface" = "SMCPolledInterface is not serializable"
  "IOPlatformSerialNumber" = "CCCCCCCCC"
  "system-type" = <02>
  "product-name" = <"MacBookPro12,1">
  "model" = <"MacBookPro12,1">
  "name" = <"/">
  "IOBusyInterest" = "IOCommand is not serializable"
}
`

func Test_extractID(t *testing.T) {
	want := "A3344D1DD-1234-22A1-B123-11AB1C11D111"
	got, err := extractID(sampleOutput)
	if err != nil {
		t.Error(err)
	}
	if got != want {
		t.Errorf("extractID() = %v, want %v", got, want)
	}
}

func Test_extractID_invalidInput(t *testing.T) {
	got, err := extractID("invalid input")
	if err == nil {
		t.Error("expected error, got none")
	}
	if got != "" {
		t.Errorf("expected empty string, got some value %s", got)
	}
	if strings.Contains(err.Error(), "Failed to extract 'IOPlatformUUID'") == false {
		t.Errorf("Got unexpected error: %v", err)
	}
}

func Test_machineID(t *testing.T) {
	got, err := machineID()
	if err != nil {
		t.Error(err)
	}
	if got == "" {
		t.Error("Got empty machine id")
	}
}
