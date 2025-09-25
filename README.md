## Social Network API

Social network API built with Go, featuring user management, posts, followers, and JWT authentication.

## Technologies

- **Go** (Golang)
- **MySQL**
- **JWT** - Auth
- **Gorilla Mux** - HTTP
- **bcrypt**
- **Insomnia** - Test

## Project Structure

```
social-network/
├── src/
│ ├── config/ # Configuration and environment variables
│ ├── controllers/ # HTTP route logic
│ ├── database/ # Database connection
│ ├── middlewares/ # Authentication and logging
│ ├── models/ # Data structs
│ ├── repositories/ # Database queries
│ ├── responses/ # Response formatters
│ ├── routes/ # Route definitions
│ └── utils/ # Utilities (hash, token, etc)
├── insomnia/ # API testing collection
├── scripts/ # SQL scripts and utilities
└── docs/ # Additional documentation
```

.env:

```
DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=
SECRET_KEY=
PORT=
```

## Install dependencies

```
go mod tidy
```

## Run the application

```
go run main.go
```

## Main Endpoints

### Users:

POST /users - Create user

GET /users - List all users

GET /users/{id} - Get user by ID

PUT /users/{id} - Update user

DELETE /users/{id} - Delete user

POST /users/{id}/follow - Follow user

DELETE /users/{id}/unfollow - Unfollow user

GET /users/{id}/followers - Get user's followers

GET /users/{id}/following - Get who user follows

### Authentication:

POST /signin - Login (get token)

POST /users/update-password - Change password

### Publications:

POST /publications - Create post

GET /publications - Get feed posts

GET /publications/{id} - Get post by ID

PUT /publications/{id} - Update post

DELETE /publications/{id} - Delete post

PUT /publications/{id}/like - Like post

PUT /publications/{id}/unlike - Unlike post

GET /users/{id}/publications - Get user's posts

## Data Model

- **users**: ID, name, nickname, email, password, created_at

- **followers**: user_id, follower_id, created_at

- **publications**: ID, title, content, author_id, likes, created_at

## Testing API

### With Insomnia:

Import the insomnia/social-network.yaml file

Configure the environment with your variables

Test all available routes

## TODO

Swagger implementation

Unit tests
