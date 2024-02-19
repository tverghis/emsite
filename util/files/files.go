package files

import (
	"errors"
	"io/fs"
	"os"
)

const uploadsDir = "uploads"

func EnsureUploadsDir() error {
	err := os.Mkdir(uploadsDir, 0755)

	if err != nil && !errors.Is(err, fs.ErrExist) {
		return err
	}

	return nil
}
