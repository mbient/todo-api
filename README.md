# TODO APP - Backend API (Go + Gin)

> [Project URL on Roadmap.sh](https://roadmap.sh/projects/todo-list-api)

This is a simple backend for a todo application with user authentication. 

## Installation

```bash
git clone https://github.com/mbient/todo-api.git
cd todo-api/
```

### Create `.env` file with following content

```bash
cat <<EOF > .env
# JWT configs
SECRET_KEY=a-secret-at-least-256-bits-long
EOF
```

## Run server

```bash
go run cmd/backend/main.go
```

## Create new user and login

```bash
# register
curl -s -X POST --json '{"Name":"John Doe", "Email":"john@doe.com", "Password":"password"}' localhost:8080/api/v1/register | jq

# login get cookie
curl -c cookiefile -X POST --json '{"Email":"john@doe.com", "Password":"password"}' localhost:8080/api/v1/login

# check protected endpoint without cookie
curl -s -X GET localhost:8080/api/v1/tasks | jq

# check endpoints with cookie
curl -b cookiefile -X GET localhost:8080/api/v1/tasks | jq

```

## API Endpoints

### Authentication
| Method | Endpoint    | Description        |
|--------|-------------|--------------------|
| POST   | `/register`   | Register new user  |
| POST   | `/login`    | Login user         |

### Tasks
| Method | Endpoint          | Description             |
|--------|-------------------|-------------------------|
| GET    | `/tasks`          | Get all tasks        |
| GET    | `/tasks/{id}`   | Get tasks by ID       |
| POST   | `/tasks`         | Add new tasks    |
| PUT   | `/tasks/{id}`         | Update tasks    |
| DELETE   | `/tasks/{id}`         | Delete tasks    |


## References:

- [IETF - JSON Web Token (JWT)](https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.2)

