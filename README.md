# go-std-markdown

[![Go Reference](https://pkg.go.dev/badge/github.com/lachlanharrisdev/go-std-markdown.svg)](https://pkg.go.dev/github.com/lachlanharrisdev/go-std-markdown)
[![Go Report Card](https://goreportcard.com/badge/github.com/lachlanharrisdev/go-std-markdown)](https://goreportcard.com/report/github.com/lachlanharrisdev/go-std-markdown)
[![codecov](https://codecov.io/gh/lachlanharrisdev/go-std-markdown/branch/main/graph/badge.svg)](https://codecov.io/gh/lachlanharrisdev/go-std-markdown)
[![GitHub license](https://img.shields.io/github/license/lachlanharrisdev/go-std-markdown.svg)](https://github.com/lachlanharrisdev/go-std-markdown/blob/main/LICENSE)

`go-std-markdown` is a go package implementing a Markdown renderer for the 
terminal.

Note: Markdown being originally designed to render as HTML, rendering in a 
terminal is occasionally challenging and some adaptation had to be made. 

Features:
- formatting
- lists
- tables
- images
- code blocks with syntax highlighting
- basic HTML support

## Usage

```go
import (
	"fmt"
	"os"

	markdown "github.com/lachlanharrisdev/go-std-markdown"
)

func main() {
	path := "README.md"
	source, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	result := markdown.Render(string(source), 80, 6)

	fmt.Println(result)
}
```

## Examples

![rendering example](misc/result.png)

![table rendering](misc/table.png)

## Origin

This is a fork of the original [go-term-markdown](https://github.com/MichaelMure/go-term-markdown)
project. This fork is primarily for maintenance purposes, and for use within
[gonetsim](https://github.com/lachlanharrisdev/gonetsim). All credit should go
to the original author.

This package has been extracted from the [git-bug](https://github.com/MichaelMure/git-bug) 
project. As such, its aim is to support this project and not to provide an 
all-in-one solution. Contributions or full-on takeover as welcome though.

## License

The MIT license of the original project has been preserved. Please see
LICENSE
