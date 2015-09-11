package which

import (
	"os"
	"strings"
)

var ext = []string{".com", ".exe", ".bat"}

func init() {
	if pathext := os.Getenv("PATHEXT"); pathext != "" {
		ext = append(strings.Split(pathext, string(os.PathListSeparator)), ext...)
	}
}
