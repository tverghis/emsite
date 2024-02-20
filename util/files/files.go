package files

import (
	"errors"
	"io"
	"io/fs"
	"mime/multipart"
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

// Save the uploaded file data to disk in the appropriate directory.
// The caller is responsible for closing the file handle.
func SaveUpload(file multipart.File) error {
	tempFile, err := os.CreateTemp("uploads", "upload-*")

	if err != nil {
		return err
	}
	defer tempFile.Close()

	if _, err := io.Copy(tempFile, file); err != nil {
		return err
	}

	return nil
}
