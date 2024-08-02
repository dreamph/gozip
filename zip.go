package gozip

import (
	"github.com/alexmullins/zip"
	"io"
	"os"
	"path/filepath"
)

func Zip(zipFileOutputPath string, filePaths []string, password ...string) error {
	if len(password) > 0 {
		return createZip(zipFileOutputPath, filePaths, password[0])
	} else {
		return createZip(zipFileOutputPath, filePaths, "")
	}
}

func createZip(zipFilePath string, filePaths []string, password string) error {
	err := zipFiles(zipFilePath, filePaths, password)
	if err != nil {
		_ = os.Remove(zipFilePath)
		return err
	}
	return nil
}

func zipFiles(zipFilePath string, filePaths []string, password string) error {
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := newZipWriter(zipFile)
	defer zipWriter.Close()

	for _, filePath := range filePaths {
		info, err := os.Stat(filePath)
		if err != nil {
			return err
		}

		if info.IsDir() {
			err := addDirectoryToZip(zipWriter, filePath, password)
			if err != nil {
				return err
			}
		} else {
			err := addFileToZip(zipWriter, filePath, filepath.Dir(filePath), password)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func addFileToZip(w *zip.Writer, file string, baseDir string, password string) error {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}

	relPath, err := filepath.Rel(baseDir, file)
	if err != nil {
		return err
	}
	header.Name = relPath

	if fileInfo.IsDir() {
		header.Name += "/"
	}

	if password != "" {
		header.SetPassword(password)
	}

	writer, err := w.CreateHeader(header)
	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		fileReader, err := os.Open(file)
		if err != nil {
			return err
		}
		defer fileReader.Close()

		_, err = io.Copy(writer, fileReader)
		if err != nil {
			return err
		}
	}

	return nil
}

func addDirectoryToZip(w *zip.Writer, dir string, password string) error {
	return filepath.Walk(dir, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return addFileToZip(w, file, dir, password)
	})
}

func newZipWriter(zipFile *os.File) *zip.Writer {
	zipWriter := zip.NewWriter(zipFile)
	return zipWriter
}
