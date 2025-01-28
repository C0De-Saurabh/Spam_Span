# Spam Span

Spam Span is a spam detection and contact management system inspired by Truecaller. It provides functionalities like spam reporting, secure contact management, and advanced search capabilities to enhance user communication.

---
## Tech Stack

### Backend
- **Golang**: A highly performant programming language used for building the backend of the application.
- **Gin Framework**: A lightweight and fast web framework for building RESTful APIs in Go.
- **GORM**: An ORM library for interacting with the SQLite database in a Go idiomatic way.

### Database
- **SQLite**: A lightweight and file-based relational database used for simplicity and quick development.

### Authentication
- **JWT (JSON Web Token)**: Used for secure authentication and session management.

### Testing & Development
- **Faker**: A library for generating fake data for testing purposes.

## Features

1. **User Registration**:  
   - Register with phone number, name, password, and optional email.
   - Example Request:
     ```json
     {
       "name": "John Doe",
       "phone": "9876543210",
       "password": "securepassword123",
       "email": "john.doe@example.com"
     }
     ```
   - Example Response:
     ```json
     {
       "message": "User registered successfully"
     }
     ```

2. **Spam Reporting**:  
   - Mark contacts as spam (daily limit enforced).
   - Example Request:
     ```json
     {
       "phone": "9123456789"
     }
     ```
   - Example Response:
     ```json
     {
       "message": "Phone number marked as spam successfully."
     }
     ```

3. **Contact Management**:  
   - Add personal contacts and synchronize with the global database.
   - View a user's personal contact list.

4. **Advanced Search**:  
   - Search by phone number or name with spam likelihood information.
   - Queries validate inputs for correctness.

5. **Testing and Development**:  
   - Random test data generation using the Faker library.

---

## Database Schema

### User Model
| Field               | Type          | Description                          |
|---------------------|---------------|--------------------------------------|
| `ID`                | `uint`        | Primary key.                         |
| `Name`              | `string`      | User's name.                         |
| `Phone`             | `string`      | Unique phone number (not null).      |
| `Password`          | `string`      | Encrypted password.                  |
| `Email`             | `string`      | User's email (optional).             |
| `Contacts`          | `[]Contact`   | One-to-many relationship with `Contact`. |
| `SpamReportsToday`  | `int`         | Count of spam reports today.         |
| `LastSpamReportDate`| `string`      | Date of last spam report.            |
| `CreatedAt`         | `time.Time`   | Record creation timestamp.           |
| `UpdatedAt`         | `time.Time`   | Record update timestamp.             |

### Contact Model
| Field         | Type          | Description              |
|---------------|---------------|--------------------------|
| `ID`          | `uint`        | Primary key.             |
| `UserID`      | `uint`        | Foreign key for User.    |
| `Name`        | `string`      | Contact's name.          |
| `Phone`       | `string`      | Contact's phone number.  |
| `Email`       | `string`      | Contact's email.         |
| `CreatedAt`   | `time.Time`   | Record creation time.    |
| `UpdatedAt`   | `time.Time`   | Record update time.      |

### GlobalContact Model
| Field         | Type          | Description              |
|---------------|---------------|--------------------------|
| `ID`          | `uint`        | Primary key.             |
| `Phone`       | `string`      | Unique phone number.     |
| `Name`        | `string`      | Default is "Unknown".    |
| `Email`       | `string`      | Email associated.        |
| `SpamReported`| `int`         | Count of spam reports.   |
| `CreatedAt`   | `time.Time`   | Record creation time.    |
| `UpdatedAt`   | `time.Time`   | Record update time.      |

---

## API Endpoints

### Authentication
- **POST /register**: Register a new user.  
- **POST /login**: Login to obtain a JWT token.  
- **POST /logout**: Log out (implementation handled in the frontend).

### Spam Reporting
- **POST /mark-spam**: Mark a phone number as spam (limited daily reports).

### Contact Management
- **POST /add-contact**: Add a new contact.  
- **GET /contacts**: View user's contact list.

### Search
- **GET /search-phone**: Search by phone number.  
- **GET /search-name**: Search by name.  

---

