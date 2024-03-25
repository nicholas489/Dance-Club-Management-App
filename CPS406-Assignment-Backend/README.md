# CPS406-Assignment-Backend
## Tech Stack
- ### GoLang
- ### GORM
- ### SQLite3

## File Tree Structure

Below is the structure of the project, explaining the purpose of each file and directory:

```plaintext
CPS406-Assignment-Backend/
├── cmd/                           # Command-line interface applications
│   └── app/                       # Main application entry point
│       └── main.go                # The main Go file where execution begins
├── internal/                      # Private application and library code
│   └── api/                       # Internal API-related code
│       └── http/                  # HTTP transport layer specific code
│           ├── coach/             # Coach-related HTTP handlers
│           │   └── coach.go       # Handles coach-related requests
│           ├── server/            # Server configuration and setup
│           │   └── router.go      # Routes and server setup
│           └── user/              # User-related HTTP handlers
│               └── user.go        # Handles user-related requests
│   └── db/                        # Database configuration and initialization
│       └── db.go                  # Database connection setup and GORM setup
│   └── util/                      # Utility functions used across the application
│       ├── sort.go                # Utility functions for sorting
│       └── util.go                # General utility functions
├── pkg/                           # Library code that's ok to use by external applications
│   ├── coach/                     # Domain model definitions for coach
│   │   └── model.go               # Coach model definition
│   ├── event/                     # Domain model definitions for event
│   │   └── model.go               # Event model definition
│   ├── jwtM/                      # Domain model definitions related to JWT middleware
│   │   └── model.go               # JWT middleware model definition
│   ├── login/                     # Domain model definitions for login functionality
│   │   └── model.go               # Login model definition
│   └── user/                      # Domain model definitions for user
│       └── model.go               # User model definition
└── .env                           # Environment variables and configuration settings
└── go.mod                         # Go module definition
└── README                         # This file

```

## Models

## Routes
| Method | Endpoint                | Description                               | Middleware                |
|--------|-------------------------|-------------------------------------------|---------------------------|
| POST   | `/login/user`           | Logs in a user                            | None                      |
| POST   | `/login/coach`          | Logs in a coach                           | None                      |
| POST   | `/signup/user`          | Signs up a new user                       | None                      |
| POST   | `/signup/coach`         | Signs up a new coach                      | None                      |
| GET    | `/users/`               | Retrieves all users                       | `JwtMiddlewareCoach`, `JwtMiddlewareAdmin` |
| GET    | `/user/{id}`            | Retrieves a specific user by ID           | `JwtMiddlewareUser`       |
| POST   | `/user/join/event`      | Allows a user to join an event            | `JwtMiddlewareUser`       |
| POST   | `/coach/event/make`     | Allows a coach to create an event         | `JwtMiddlewareCoach`, `JwtMiddlewareAdmin` |
| GET    | `/coach/{name}`         | Retrieves a specific event by name        | `JwtMiddlewareCoach`, `JwtMiddlewareAdmin` |



