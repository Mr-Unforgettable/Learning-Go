# Address Book

The Address Book project is a simple command-line application built in Go that allows users to manage their contacts. It provides basic functionality to add, view, search, update, and delete contacts. Contacts are stored in memory and can be saved to a file for persistence.

## Features

The Address Book application includes the following features:

- Add a Contact: Allows the user to add a new contact by providing their name, email, and phone number.
- View Contacts: Displays a list of all contacts in the address book, showing their names and basic information.
- Search Contacts: Provides a search functionality to find a specific contact by name or other criteria.
- Update Contact: Allows the user to update the information of an existing contact, such as modifying the email or phone number.
- Delete Contact: Provides an option to delete a contact from the address book.
- Save and Load Contacts: Supports saving contacts to a file for persistence, and loading them again when the application is restarted.
- Input Validation: Validates user input to ensure it meets the required format or constraints.

## Folder Structure

The project follows the following folder structure:

```
address-book/
    ├── cmd/
    │   └── main.go
    ├── internal/
    │   ├── addressbook/
    │   │   ├── addressbook.go
    │   │   └── contacts.go
    │   └── filehandler/
    │       └── filehandler.go
    ├── pkg/
    │   └── utils/
    │       └── input.go
    └── README.md
```

- `cmd/`: Contains the `main.go` file, which serves as the entry point of the application.
- `internal/addressbook/`: Implements the address book logic, including the address book struct, contact operations, and search functionality.
- `internal/filehandler/`: Handles file operations for saving and loading contacts.
- `pkg/utils/`: Provides utility functions for input validation and user interaction.
- `README.md`: Contains project documentation and instructions.

## Getting Started

To run the Address Book application on your local machine, follow these steps:

1. Make sure you have Go installed on your system. You can download and install Go from the official website: [https://golang.org/](https://golang.org/)

2. Clone this repository to your local machine or download the source code as a ZIP file.

3. Navigate to the project's root directory.

4. Build the application by running the following command:

   ```shell
   go build -o address-book ./cmd
   ```

5. Run the application:

   ```shell
   ./address-book
   ```

6. Follow the on-screen instructions to interact with the Address Book.

## Contributing

Contributions to this project are welcome! If you find any issues or have ideas for enhancements, feel free to open an issue or submit a pull request. Please ensure that your contributions align with the project's coding style and conventions.

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute the code as permitted by the license.

## Acknowledgments

- The project was inspired by the need for a simple address book management system.
- Special thanks to the Go community for their excellent resources and libraries.

Feel free to customize the README based on your project's specific requirements and add any additional sections or information that you find relevant.