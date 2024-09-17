package util

import (
	"os"
	"strings"

	"github.com/alexmullins/zip"
)

func CreateEncryptedZip(name, content, password string, tags ...string) error {
	// Create the zip file
	zipFile, err := os.Create(name + ".zip")
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// Initialize zip writer with encryption support
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// create file name from tags and content
	fileName := strings.Join(append(tags, "content"), "_") + ".txt"

	// Create a new entry for content.txt in the zip file
	fileWriter, err := zipWriter.Encrypt(fileName, password)
	if err != nil {
		return err
	}

	// Write the content to the entry
	_, err = fileWriter.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}
