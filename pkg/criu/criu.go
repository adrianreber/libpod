package criu

import (
	"github.com/checkpoint-restore/criu/lib/go/src/criu"
)

// CheckForCriu uses CRIU's go bindings to check if the CRIU
// binary exists and if it at least the version Podman needs.
func CheckForCriu() bool {
	c := criu.MakeCriu()
	// For Podman at least CRIU 3.11 is required
	result, err := c.IsCriuAtLeast(31100)
	if err != nil {
		return false
	}
	return result
}
