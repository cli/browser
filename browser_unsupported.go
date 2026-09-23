//go:build !linux && !windows && !darwin && !openbsd && !freebsd && !netbsd
// +build !linux,!windows,!darwin,!openbsd,!freebsd,!netbsd

package browser

import (
	"context"
	"fmt"
	"runtime"
)

func openBrowser(_ context.Context, _ string) error {
	return fmt.Errorf("openBrowser: unsupported operating system: %v", runtime.GOOS)
}
