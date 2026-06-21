package browser

import (
	"context"

	"golang.org/x/sys/windows"
)

func openBrowser(_ context.Context, url string) error {
	return windows.ShellExecute(0, nil, windows.StringToUTF16Ptr(url), nil, nil, windows.SW_SHOWNORMAL)
}
