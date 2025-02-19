package browser

import (
	"os/exec"
	"strings"
)

// providersForUrl returns a slice of possible providers to open a browser on linux. The
// order is from most to least preferred.
func providersForUrl(url string) []string {
	base := []string{"xdg-open", "x-www-browser", "www-browser"}
	// explorer.exe is available in WSL. It can open http:// and https://, but not
	// file://. explorer.exe is preferred over wslview for performance and to avoid
	// redundant requests.
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		base = append(base, "explorer.exe", "wslview")
	} else {
		base = append(base, "wslview")
	}
	return base
}

func openBrowser(url string) error {
	providers := providersForUrl(url)

	// There are multiple possible providers to open a browser on linux
	// One of them is xdg-open, another is x-www-browser, then there's www-browser, etc.
	// Look for one that exists and run it
	for _, provider := range providers {
		if _, err := exec.LookPath(provider); err == nil {
			return runCmd(provider, url)
		}
	}

	return &exec.Error{Name: strings.Join(providers, ","), Err: exec.ErrNotFound}
}
