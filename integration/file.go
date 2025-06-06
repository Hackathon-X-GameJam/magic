package integration

import (
	"io"
	"os"
)

func CopyFile(source string, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}
	return nil
}

// Create a new file with content.
func CreateFileWithContent(name string, content string) error {
	return os.WriteFile(name, []byte(content), 0755)
}
