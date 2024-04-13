package fileutils

import(
	"os"
)

func ReadFile(path string) ([]byte, error) {

	data, err := os.ReadFile(path)
	return data, err
}
