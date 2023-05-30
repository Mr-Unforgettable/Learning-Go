package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"internal/addressbook"
)

func main() {
	// Create a new address book instance.
	book := addressbook.NewAddressBook()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("************************")
	fmt.Println("Welcome to Address Book!")
	fmt.Println("************************")

	for {
		fmt.Println("\nPlease choose an option:")
		fmt.Println("1. Add a contact")
		fmt.Println("2. View contacts")
		fmt.Println("3. Search contacts")
		fmt.Println("4. Update a contact")
		fmt.Println("5. Delete a contact")
		fmt.Println("6. Save contacts")
		fmt.Println("0. Exit")

		fmt.Println("============================")
		fmt.Print("Option: ")
		option, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Println("\nAdd a Contact")
			fmt.Println("===========================")

			contact := getAddressFromUser(reader)
			book.AddContact(&contact)

			fmt.Println("===========================")
			fmt.Println("Contact added Successfully.")
			fmt.Println("===========================")

		case "2":
			fmt.Println("\nView contacts")
			fmt.Println("===========================")

			contacts := book.GetAllContacts()
			if len(contacts) == 0 {
				fmt.Println("No contacts found.")
			} else {
				fmt.Println("Contacts:")
				for _, contact := range contacts {
					fmt.Println("=========================")
					fmt.Printf("-Name: %s\n -Email: %s\n -Phone: %s\n", contact.Name, contact.Email, contact.Phone)
					fmt.Println("=========================")
				}
			}
		case "3":
			fmt.Println("\nSearch Contact")
			fmt.Println("===========================")

			fmt.Print("Enter a name or keyword to search: ")
			keyword, _ := reader.ReadString('\n')
			keyword = strings.TrimSpace(keyword)

			contacts := book.SearchContacts(keyword)
			if len(contacts) == 0 {
				fmt.Println("No contacts found.")
			} else {
				fmt.Println("Matching contacts: ")
				for _, contact := range contacts {
					fmt.Println("=========================")
					fmt.Printf("-Name: %s\n -Email: %s\n -Phone: %s\n", contact.Name, contact.Email, contact.Phone)
					fmt.Println("=========================")
				}
			}

		case "4":
			fmt.Println("\nUpdate a contact")
			fmt.Println("=========================")

			fmt.Print("Enter the name of the contact to update: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			contact := book.GetContactByName(name)
			if contact == nil {
				fmt.Println("Contact not found.")
			} else {
				fmt.Printf("Current contact: %s\n", contact.Name)
				fmt.Println("Enter new details: ")

				newContact := getAddressFromUser(reader)
				book.UpdateContact(contact, &newContact)

				fmt.Println("=============================")
				fmt.Println("Contact updated successfully.")
				fmt.Println("=============================")
			}

		case "5":
			fmt.Println("\nDelete a contact")
			fmt.Println("==============================")

			fmt.Print("Enter the name of the contact to delete: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			contact := book.GetContactByName(name)
			if contact == nil {
				fmt.Println("Contact not found.")
			} else {
				book.DeleteContact(contact)
				fmt.Println("=============================")
				fmt.Println("Contact deleted successfully.")
				fmt.Println("=============================")
			}

		case "6":
			fmt.Println("\n Save Contacts")
			fmt.Println("=================================")

			fmt.Print("Enter the filepath to save contacts: ")
			filePath, _ := reader.ReadString('\n')
			filePath = strings.TrimSpace(filePath)

			err := book.SaveContactsToFile(filePath)
			if err != nil {
				fmt.Printf("Failed to save contacts: %v\n", err)
			} else {
				fmt.Println("============================")
				fmt.Println("Contacts saved successfully.")
				fmt.Println("============================")
			}

		case "0":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid option, Please try again.")
		}
	}
}

func getAddressFromUser(reader *bufio.Reader) addressbook.Contact {
	fmt.Println("Enter name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Enter phone number: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	return addressbook.Contact{
		Name:  name,
		Email: email,
		Phone: phone,
	}
}
