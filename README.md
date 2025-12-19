This is a RESTful API built with Go to manage users with their Name and Date of Birth (DOB). The application dynamically calculates a user's age when fetching details.

🛠 Tech Stack

    Framework: GoFiber 

    Database: PostgreSQL

    DB Access: SQLC (Type-safe generated code)

    Logging: Uber Zap

    Validation: go-playground/validator

📂 Project Structure

    /cmd/server/main.go     # Application entry point
    /config/                # DB connection and env loading
    /db/migrations/         # SQL schema migrations
    /db/sqlc/               # Generated SQLC code
    /internal/              # Core logic
    ├── handler/          # HTTP handlers
    ├── repository/       # Database access layer
    ├── service/          # Business logic (Age calculation)
    ├── routes/           # API route definitions
    ├── middleware/       # Custom middlewares (RequestID, Logger)
    ├── models/           # Request/Response structs
    └── logger/           # Uber Zap initialization


Setup Instructions

1. Database Setup

Ensure you have PostgreSQL installed and running on your machine.

    1. Log into PostgreSQL: sudo -u postgres psql

    2. Create the database: CREATE DATABASE user_db;

    3. Connect to the DB: \c user_db

    4. Create the table:

    1. Database Setup

Ensure you have PostgreSQL installed and running on your machine.

    Log into PostgreSQL: sudo -u postgres psql

    Create the database: CREATE DATABASE user_db;

    Connect to the DB: \c user_db

    Create the table:
        CREATE TABLE users (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL,
            dob DATE NOT NULL
        );

2. Environment Configuration

Create a .env file in the root directory:

    DB_USER=
    DB_PASSWORD=
    DB_NAME=
    DB_HOST=
    DB_PORT=

3. Install Dependencies

    go mod tidy

4. Run the Application

    go run cmd/server/main.go

📡 API Endpoints

    Method,Endpoint,Description
    POST,/users,Create a new user 
    GET,/users/:id,Get user details (with dynamic age) 
    GET,/users,List all users 
    PUT,/users/:id,Update user details 
    DELETE,/users/:id,Delete a user 