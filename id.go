// Package machineid provides support for reading the unique machine id of most host OS's (without admin privileges).
// https://github.com/denisbrodbeck/machineid
// Supported OS systems: BSD, Linux, OS X, Windows
package machineid // import "github.com/denisbrodbeck/machineid"

// ID returns the platform specific machine id of the current host OS.
// Regard the returned id as "confidential" and consider using ProtectedID() instead.
func ID() (string, error) {
	return machineID()
}

// ProtectedID returns a hashed version of the machine ID in a cryptographically secure way,
// using a fixed, application-specific key.
// Internally, this function calculates HMAC-SHA256 of the application ID, keyed by the machine ID.
func ProtectedID(appID string) (string, error) {
	id, err := ID()
	if err != nil {
		return "", err
	}
	return protect(appID, id), nil
}
