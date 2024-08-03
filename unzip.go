package gozip

import (
	"github.com/alexmullins/zip"
	"io"
	"os"
	"path/filepath"
)

func Unzip(source string, destination string, password ...string) error {
	if len(password) > 0 {
		return unzip(source, destination, password[0])
	} else {
		return unzip(source, destination, "")
	}
}

func unzip(source, destination string, password string) error {
	archive, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer archive.Close()
	for _, file := range archive.Reader.File {
		if password != "" {
			if file.IsEncrypted() {
				file.SetPassword(password)
			}
		}

		if file.FileInfo().IsDir() {
			_ = os.MkdirAll(filepath.Join(destination, file.Name), os.ModePerm)
			continue
		}
		reader, err := file.Open()
		if err != nil {
			return err
		}
		defer reader.Close()
		path := filepath.Join(destination, file.Name)
		// Remove file if it already exists; no problem if it doesn't; other cases can error out below
		_ = os.Remove(path)
		// Create a directory at path, including parents
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
		// If file is _supposed_ to be a directory, we're done
		if file.FileInfo().IsDir() {
			continue
		}
		// otherwise, remove that directory (_not_ including parents)
		err = os.Remove(path)
		if err != nil {
			return err
		}
		// and create the actual file.  This ensures that the parent directories exist!
		// An archive may have a single file with a nested path, rather than a file for each parent dir
		writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer writer.Close()
		_, err = io.Copy(writer, reader)
		if err != nil {
			return err
		}
	}
	return nil
}
