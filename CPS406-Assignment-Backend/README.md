# CPS406-Assignment-Backend
## Tech Stack
- ### GoLang
- ### GORM
- ### SQLite3
- ### JWT

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

## Data Models

The following tables describe the data models used in the application and their respective fields:

### Coach

| Field       | Type   | JSON Key       | GORM Annotation | Description                              |
|-------------|--------|----------------|-----------------|------------------------------------------|
| ID          | uint   | -              | -               | The primary key (auto-incremented).      |
| CreatedAt   | time   | -              | -               | Timestamp of creation.                   |
| UpdatedAt   | time   | -              | -               | Timestamp of last update.                |
| DeletedAt   | time   | -              | -               | Timestamp of deletion (if soft deleted). |
| Name        | string | `name`         | -               | Unique username for the coach.           |
| Email       | string | `email`        | `index;unique`  | Unique email address for the coach.      |
| PhoneNumber | int    | `phone_number` | -               | Contact phone number for the coach.      |
| Password    | string | `password`     | -               | Hashed password for the coach's account. |

### Event

| Field      | Type        | JSON Key      | GORM Annotation | Description                              |
|------------|-------------|---------------|-----------------|------------------------------------------|
| ID         | uint        | -             | -               | The primary key (auto-incremented).      |
| CreatedAt  | time        | -             | -               | Timestamp of creation.                   |
| UpdatedAt  | time        | -             | -               | Timestamp of last update.                |
| DeletedAt  | time        | -             | -               | Timestamp of deletion (if soft deleted). |
| Name       | string      | `name`        | `index;unique`  | Unique name for the event.               |
| CoachEmail | string      | `coach_email` | -               | Email of the coach hosting the event.    |
| Location   | string      | `location`    | -               | Location where the event will be held.   |
| Cost       | int         | `cost`        | -               | Cost to attend the event.                |
| Users      | []user.User | `users`       | -               | List of users attending the event.       |

### User

| Field       | Type   | JSON Key       | GORM Annotation | Description                              |
|-------------|--------|----------------|-----------------|------------------------------------------|
| ID          | uint   | -              | -               | The primary key (auto-incremented).      |
| CreatedAt   | time   | -              | -               | Timestamp of creation.                   |
| UpdatedAt   | time   | -              | -               | Timestamp of last update.                |
| DeletedAt   | time   | -              | -               | Timestamp of deletion (if soft deleted). |
| Name        | string | `name`         | -               | Full name of the user.                   |
| Password    | string | `password`     | -               | Hashed password for the user's account.  |
| Email       | string | `email`        | `index;unique`  | Unique email address for the user.       |
| PhoneNumber | int    | `phone_number` | -               | Contact phone number for the user.       |
| Balance     | int    | `balance`      | -               | Balance amount for the user's account.   |
| EventID     | uint   | `event_id`     | -               | Foreign key relating to an Event.        |

## Seeded Data for Testing

Below are tables with sample data that is inserted into the database by the seeder function for testing purposes.

### Coaches

| UserName  | Email                | PhoneNumber | Password   |
|-----------|----------------------|-------------|------------|
| CoachMike | mike@example.com     | 1234567890  | pass123    |
| CoachAnna | anna@example.com     | 1234567891  | pass456    |

### Events

| Name          | Coach Email       | Location      | Cost |
|---------------|-------------------|---------------|------|
| Morning Yoga  | mike@example.com  | Central Park  | 10   |
| Evening Run   | anna@example.com  | Riverside     | 5    |

### Users

| Name       | Email                  | PhoneNumber | Password  | Balance |
|------------|------------------------|-------------|-----------|---------|
| John Doe   | john.doe@example.com   | 1234567892  | secure123 | 100     |
| Jane Smith | jane.smith@example.com | 1234567893  | secure456 | 150     |

This data is meant for initial development and testing only and should not be used in production environments.


## Routes
| Method | Endpoint            | Description                        | Middleware                                 |
|--------|---------------------|------------------------------------|--------------------------------------------|
| POST   | `/login/user`       | Logs in a user                     | None                                       |
| POST   | `/login/coach`      | Logs in a coach                    | None                                       |
| POST   | `/signup/user`      | Signs up a new user                | None                                       |
| POST   | `/signup/coach`     | Signs up a new coach               | None                                       |
| GET    | `/users/`           | Retrieves all users                | `JwtMiddlewareCoach`, `JwtMiddlewareAdmin` |
| GET    | `/user/{id}`        | Retrieves a specific user by ID    | `JwtMiddlewareUser`                        |
| POST   | `/user/join/event`  | Allows a user to join an event     | `JwtMiddlewareUser`                        |
| POST   | `/coach/event/make` | Allows a coach to create an event  | `JwtMiddlewareCoach`, `JwtMiddlewareAdmin` |
| GET    | `/coach/{name}`     | Retrieves a specific event by name | `JwtMiddlewareCoach`, `JwtMiddlewareAdmin` |



