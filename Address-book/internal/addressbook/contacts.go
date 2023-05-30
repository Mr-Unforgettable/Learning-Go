package addressbook

// NewContact creates a new contact in the addressBook.
func NewContact(name, email, phone string) *Contact {
	return &Contact{
		Name:  name,
		Email: email,
		Phone: phone,
	}
}

// GetName returns the name of the contact.
func (c *Contact) GetName() string {
	return c.Name
}

// GetEmail returns the email of the contact.
func (c *Contact) GetEmail() string {
	return c.Email
}

// GetPhone returns the phone of the contact.
func (c *Contact) GetPhone() string {
	return c.Phone
}

// SetName sets the name of the contact.
func (c *Contact) SetName(name string) {
	c.Name = name
}

// SetEmail sets the Email of the contact.
func (c *Contact) SetEmail(email string) {
	c.Email = email
}

// SetPhone sets the Phone of the contact.
func (c *Contact) SetPhone(phone string) {
	c.Phone = phone
}
