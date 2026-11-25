package browser

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
)

func openBrowser(ctx context.Context, url string) error {
	err := runCmd(ctx, "xdg-open", url)
	if errors.Is(err, exec.ErrNotFound) {
		return fmt.Errorf("%w - install xdg-utils from ports(8)", err)
	}
	return err
}
