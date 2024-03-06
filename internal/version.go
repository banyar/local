package internal

import (
	"fmt"

	semver "github.com/blang/semver/v4"
)

var (
	// Cmd version. This var must be injected using Go -ldflags during build time.
	version = "0.0.1"
)

// Get the sem-ver compatible app version.
// The actual version value is supposed to be injected with -ldflags during build time.
func Version() string {
	if v, err := semver.Make(version); err != nil {
		return fmt.Sprintf("Parsing version err: %v", err)
	} else {
		return fmt.Sprint(v.String())
	}
}
