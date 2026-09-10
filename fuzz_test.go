package markdown

import (
	"testing"
)

func FuzzRender(f *testing.F) {
	seeds := []string{
		"# Heading\n\nSome **bold** and `code`.\n",
		"```go\npackage main\nfunc main() {}\n```\n",
		"| a | b |\n|---|----|\n| 1 | 2  |\n",
		"> blockquote\n>\n> with *emphasis*\n",
		"1. item one\n2. item two\n",
		"<span><a href=\"https://example.com\">link</a></span>\n",
		"~~strike~~ and [inline](https://example.com)\n",
		"",
	}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, source string) {
		_ = Render(source, 80, 6)
	})
}
