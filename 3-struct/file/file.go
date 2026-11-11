package file

import (
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func IsJSON(path string) bool {
	ext := filepath.Ext(path)
	return strings.ToLower(ext) == ".json"
}
