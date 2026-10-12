package browser

import "context"

func openBrowser(ctx context.Context, url string) error {
	return runCmd(ctx, "open", url)
}
