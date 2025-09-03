# 💚 User API

Welcome to my User API, a simple Go application that lets you manage users by creating, updating, deleting, and listing them in memory.

## 🚀 Features
-  Create users with name and email

-  Retrieve a single user by ID

-  Update user information

-  Delete users

- Health check endpoint to verify the API is running

## 🛠 Installation & Usage
```
# Clone this repository
git clone https://github.com/tasteofsun/user-api.git

# Enter the project directory
cd user-api

# Install dependencies
go mod tidy

# Run the API
go run main.go


The API will be available at:
http://localhost:8080
```

## 💻 Testing the API
You can use Postman or any API client to test endpoints at http://localhost:8080/api/users.

Examples using cURL:

# Health Check
```
curl -X GET http://localhost:8080/health

{ "status": "OK", "service": "user-api" }
```

# Get all users
```
curl -X GET http://localhost:8080/api/users

{
  "users": [
    { "id": "uuid", "name": "John Doe", "email": "john@example.com" }
  ],
  "total": 1
}
```

# Get user by ID
```
curl -X GET http://localhost:8080/api/users/<USER_ID>

{ "id": "uuid", "name": "John Doe", "email": "john@example.com" }
```

# Create a new user
```
curl -X POST http://localhost:8080/api/users \
-H "Content-Type: application/json" \
-d '{"name": "John Doe", "email": "john@example.com"}'

{ "id": "uuid", "name": "John Doe", "email": "john@example.com" }
```

# Update user
```
curl -X PUT http://localhost:8080/api/users/<USER_ID> \
-H "Content-Type: application/json" \
-d '{"name": "Jane Doe", "email": "jane@example.com"}'

{ "id": "uuid", "name": "Jane Doe", "email": "jane@example.com" }
```

# Delete user
```
curl -X DELETE http://localhost:8080/api/users/<USER_ID>

{ "message": "User deleted successfully" }
```

📂 Project Structure
```
user-api/
│── main.go
│── handlers/
│   └── user_handlers.go
│── models/
│   └── user.go
│── go.mod
│── README.md
```

## 🖥️ Technologies Used

- 🐹 Go

- ⚡ Gin Gonic framework

- 🆔 Google UUID for unique IDs

- 📦 JSON for request/response

- 💼 Postman (for testing API endpoints)

## 🤝 Contributions
Feel free to fork this project, submit pull requests, or open issues for improvements! 🚀
