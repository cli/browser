// Package browser provides helpers to open files, readers, and urls in a browser window.
//
// The choice of which browser is started is entirely client dependent.
package browser

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

// Stdout is the io.Writer to which executed commands write standard output.
var Stdout io.Writer = os.Stdout

// Stderr is the io.Writer to which executed commands write standard error.
var Stderr io.Writer = os.Stderr

// OpenFile opens new browser window for the file path.
func OpenFile(path string) error {
	return OpenFileContext(context.Background(), path)
}

// OpenFileContext opens new browser window for the file path.
// Accepts a context.Context to allow for cancelling the operation.
func OpenFileContext(ctx context.Context, path string) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	return OpenURLContext(ctx, "file://"+path)
}

// OpenReader consumes the contents of r and presents the
// results in a new browser window.
func OpenReader(r io.Reader) error {
	return OpenReaderContext(context.Background(), r)
}

// OpenReaderContext consumes the contents of r and presents the
// results in a new browser window.
// Accepts a context.Context to allow for cancelling the operation.
func OpenReaderContext(ctx context.Context, r io.Reader) error {
	f, err := os.CreateTemp("", "browser.*.html")
	if err != nil {
		return fmt.Errorf("browser: could not create temporary file: %w", err)
	}
	if _, err := io.Copy(f, r); err != nil {
		f.Close()
		return fmt.Errorf("browser: caching temporary file failed: %w", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("browser: caching temporary file failed: %w", err)
	}
	return OpenFileContext(ctx, f.Name())
}

// OpenURL opens a new browser window pointing to url.
func OpenURL(url string) error {
	return OpenURLContext(context.Background(), url)
}

// OpenURLContext opens a new browser window pointing to url.
// Accepts a context.Context to allow for cancelling the operation.
func OpenURLContext(ctx context.Context, url string) error {
	return openBrowser(ctx, url)
}

func runCmd(ctx context.Context, prog string, args ...string) error {
	if ctx == nil {
		ctx = context.Background()
	}
	cmd := exec.CommandContext(ctx, prog, args...)
	cmd.Stdout = Stdout
	cmd.Stderr = Stderr
	return cmd.Run()
}
