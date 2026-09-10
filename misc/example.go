package misc

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
