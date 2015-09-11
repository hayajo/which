package which

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Which(exec string) (string, error) {
	if exec == "" {
		return exec, errors.New("empty string")
	}

	var found string

	for _, path := range getPath() {
		for _, ext := range getExt() {
			file := filepath.Join(path, exec+ext)

			info, err := os.Stat(file)
			if err != nil {
				if !os.IsNotExist(err) {
					return file, err
				}
				continue
			}

			mode := info.Mode()

			if !mode.IsDir() && mode.Perm()&os.FileMode(perm) != 0 {
				found = file
				break
			}
		}

		if found != "" {
			break
		}
	}

	if found == "" {
		return found, errors.New(exec + " not found")
	}

	return found, nil
}

func getPath() []string {
	envPath := os.Getenv("PATH")
	curDir, _ := os.Getwd()
	path := strings.Split(envPath, string(os.PathListSeparator))
	path = append([]string{curDir}, path...)
	return path
}

func getExt() []string {
	return ext
}
