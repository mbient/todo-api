# TODO APP - Backend API (Go + Gin)

> [Project URL on Roadmap.sh](https://roadmap.sh/projects/todo-list-api)

This is a simple backend for a todo application with user authentication. 

## Installation

```bash
git clone https://github.com/mbient/todo-api.git
cd todo-api/backend/
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
go run main.go
```

## Create new user and login

```bash
# register
curl -s -X POST \
--json '{"Name":"John Doe", "Email":"john@doe.com", "Password":"password"}' \
localhost:8080/api/v1/register | jq

# login and get token
TOKEN=$(curl -s -X POST \
--json '{"Email":"john@doe.com", "Password":"password"}' \
localhost:8080/api/v1/login | jq -r '.token')

# get protected site
curl -s -X GET \
-H "Authorization: Bearer $TOKEN" \
localhost:8080/api/v1/tasks | jq

curl -s -X POST \
-H "Authorization: Bearer $TOKEN" \
--json '{"Title":"Filtering and sorting", "Description":"Implement filtering and sorting for the todo-api"}' \
localhost:8080/api/v1/tasks | jq

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

