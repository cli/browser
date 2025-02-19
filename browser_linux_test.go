package browser

import (
	"slices"
	"testing"
)

func TestProvidersForUrl(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want []string
	}{
		{
			name: "http protocol",
			url:  "http://example.com",
			want: []string{"xdg-open", "x-www-browser", "www-browser", "explorer.exe", "wslview"},
		},
		{
			name: "https protocol",
			url:  "https://example.com",
			want: []string{"xdg-open", "x-www-browser", "www-browser", "explorer.exe", "wslview"},
		},
		{
			name: "file protocol",
			url:  "file:///path/to/file",
			want: []string{"xdg-open", "x-www-browser", "www-browser", "wslview"},
		},
		{
			name: "no protocol",
			url:  "example.sh",
			want: []string{"xdg-open", "x-www-browser", "www-browser", "wslview"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := providersForUrl(tt.url)
			if !slices.Equal(got, tt.want) {
				t.Errorf("providersForUrl(%q) = %v; want %v", tt.url, got, tt.want)
			}
		})
	}
}
