//go:build go1.24

package html_test

import (
	"io"
	"testing"

	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func BenchmarkBooleanAttributes(b *testing.B) {
	booleanAttributes := []func() g.Node{
		html.Async,
		html.AutoFocus,
		html.AutoPlay,
		html.Checked,
		html.Controls,
		html.Defer,
		html.Disabled,
		html.Loop,
		html.Multiple,
		html.Muted,
		html.Open,
		html.PlaysInline,
		html.ReadOnly,
		html.Required,
		html.Selected,
		html.FormNoValidate,
		html.Inert,
	}

	for b.Loop() {
		for _, attr := range booleanAttributes {
			attr().Render(io.Discard)
		}
	}
}
