package files

import (
	"archive/tar"
	"bytes"
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

// Return tarball of all uploaded files so far.
func GetUploadsArchive() (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)

	tw := tar.NewWriter(buf)
	tw.AddFS(os.DirFS(uploadsDir))

	if err := tw.Close(); err != nil {
		return nil, err
	}

	return buf, nil
}
