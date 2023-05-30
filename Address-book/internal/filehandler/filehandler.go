package filehandler

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

// Contact represents contacts in the addressBook.
type Contact struct {
	Name  string
	Email string
	Phone string
}

// FileHandler handles reading a writing contacts to a file.
type FileHandler struct {
	filePath string
}

// NewFileHandler creates a new instance of the FileHandler with the filePath.
func NewFileHandler(filePath string) *FileHandler {
	return &FileHandler{
		filePath: filePath,
	}
}

// WriteFile writes data to a file using bufio.Writer.
func WriteFile(filePath string, data []byte, perm os.FileMode) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if _, err := writer.Write(data); err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush writer: %v", err)
	}
	return nil
}

// ReadContactsFromFile reads contacts from file and returns them as a slice of contacts.
func (fh *FileHandler) ReadContactsFromFile() ([]*Contact, error) {
	file, err := os.Open(fh.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var contacts []*Contact
	for _, record := range records {
		if len(record) == 3 {
			contact := &Contact{
				Name:  record[0],
				Email: record[1],
				Phone: record[2],
			}
			contacts = append(contacts, contact)
		}
	}
	return contacts, nil
}

// WriteContactsToFile writes the contacts to the file.
func (fh *FileHandler) WriteContactsToFile(contacts []*Contact) error {
	file, err := os.Create(fh.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, contact := range contacts {
		record := []string{
			contact.Name,
			contact.Email,
			contact.Phone,
		}
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}
	return nil
}
