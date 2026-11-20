package browser

import (
	"context"
	"errors"
	"os/exec"
)

func openBrowser(ctx context.Context, url string) error {
	err := runCmd(ctx, "xdg-open", url)
	if e, ok := err.(*exec.Error); ok && e.Err == exec.ErrNotFound {
		return errors.New("xdg-open: command not found - install xdg-utils from pkgsrc(7)")
	}
	return err
}
