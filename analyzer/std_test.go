package analyzer

import "testing"

func Test_isStdLib(t *testing.T) {
	tests := map[string]bool{
		"context.Context":                     true,
		"io/fs.File":                          true,
		"github.com/user/pkg/context.Context": false,
		"foo/bar.Context":                     false,
	}

	for name, want := range tests {
		want, name := want, name
		t.Run(name, func(t *testing.T) {
			got := isStdLib(name)
			assert(t, want == got,
				"pkg %s doens't match expectations (got %v vs want %v)", name, got, want)
		})
	}
}