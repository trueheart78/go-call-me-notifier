package version

import (
	"fmt"
	"runtime"
)

const (
	// Library name
	Library = "go-call-me-notifier"
	// Current version of the library
	Current = 1.0

	// ReleaseCandidate version of the library
	ReleaseCandidate = 0
)

// Full returns the full version string
func Full() string {
	var rc string
	if ReleaseCandidate > 0 {
		rc = fmt.Sprintf("-rc%d", ReleaseCandidate)
	}
	return fmt.Sprintf("%v version %.1f%v %v/%v", Library, Current, rc, runtime.GOOS, runtime.GOARCH)
}
