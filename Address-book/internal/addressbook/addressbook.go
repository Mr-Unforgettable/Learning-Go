package addressbook

import (
	"encoding/json"
	"filehandler"
	"strings"
)

// AddressBook represents the address book.
type AddressBook struct {
	contacts []*Contact
}

// Contact represents a contact in the address book.
type Contact struct {
	Name  string
	Email string
	Phone string
}

// NewAddressBook creates a new instance of AddressBook.
func NewAddressBook() *AddressBook {
	return &AddressBook{
		contacts: []*Contact{},
	}
}

// AddContact adds a new contact to the address.
func (ab *AddressBook) AddContact(contact *Contact) {
	ab.contacts = append(ab.contacts, contact)
}

// GetAllContacts returns all contacts in the addressBook.
func (ab *AddressBook) GetAllContacts() []*Contact {
	return ab.contacts
}

// GetContactByName returns the contact with the give name.
func (ab *AddressBook) GetContactByName(name string) *Contact {
	for _, contact := range ab.contacts {
		if contact.Name == name {
			return contact
		}
	}
	return nil
}

// SearchContacts searchs the contact in the address book based on a keyword.
func (ab *AddressBook) SearchContacts(keyword string) []*Contact {
	var results []*Contact

	for _, contact := range ab.contacts {
		if containsKeyword(contact, keyword) {
			results = append(results, contact)
		}
	}
	return results
}

// UpdateContact updates the existing contact with new details.
func (ab *AddressBook) UpdateContact(contact *Contact, newContact *Contact) {
	contact.Name = newContact.Name
	contact.Email = newContact.Email
	contact.Phone = newContact.Phone
}

// DeleteContact deletes the contact from the address book.
func (ab *AddressBook) DeleteContact(contact *Contact) {
	for i, c := range ab.contacts {
		if c == contact {
			// Remove the contact from the slice
			ab.contacts = append(ab.contacts[:i], ab.contacts[i+1:]...)
			return
		}
	}
}

// SaveContactsToFile saves the contacts to a file at the specified filePath.
func (ab *AddressBook) SaveContactsToFile(filePath string) error {
	// Converts the contacts to JSON format
	contactsJSON, err := json.Marshal(ab.contacts)
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	// err = bufio.WriteFile(filePath, contactsJSON, 0644)
	err = filehandler.WriteFile(filePath, contactsJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}

// containsKeyword check if the contact's name, email & phone contains that keyword.
func containsKeyword(contact *Contact, keyword string) bool {
	keyword = strings.TrimSpace(keyword)
	lowerName := strings.ToLower(contact.Name)
	lowerEmail := strings.ToLower(contact.Email)
	lowerPhone := strings.ToLower(contact.Phone)

	return strings.Contains(lowerName, keyword) || strings.Contains(lowerEmail, keyword) || strings.Contains(lowerPhone, keyword)
}
