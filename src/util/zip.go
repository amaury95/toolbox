package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexmullins/zip"
)

func CreateEncryptedZip(name, content, password string, tags ...string) error {
	// Remove .zip extension if it exists, it will be added later
	name = strings.TrimSuffix(name, ".zip")

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

func ReadEncryptedZip(file, password string, cleanOutput bool) error {
	// get absolute path of file
	absPath, err := filepath.Abs(file)
	if err != nil {
		return err
	}

	zipReader, err := zip.OpenReader(absPath)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, file := range zipReader.File {
		if file.IsEncrypted() {
			file.SetPassword(password)
			f, err := file.Open()
			if err != nil {
				log.Fatal(err)
			}
			content, err := io.ReadAll(f)
			if err != nil {
				log.Fatal(err)
			}
			if cleanOutput {
				fmt.Println(string(content))
			} else {
				log.Printf("%s: %s", file.Name, string(content))
			}
		}
	}

	return nil
}
