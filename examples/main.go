package main

import (
	"fmt"
	"os"

	"github.com/joaosoft/color"
)

func main() {
	fmt.Fprintf(os.Stdout, fmt.Sprintf("%s joao", color.WithColor("hello", color.FormatBold, color.ForegroundRed, color.BackgroundCyan)))
}
