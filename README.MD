# GoAuth service

GoAuth service is a template for authenticating service written in [Golang](https://go.dev/). With the support of [Gin](https://github.com/gin-gonic/gin), the service is implemented based on Clean Architecture. All the verifying process is powered by [JWT](https://pkg.go.dev/github.com/golang-jwt/jwt/v5).

**Key features**:
- HTTP Basic Authentication for getting JWT
- Create a new account
- Get list of created accounts
- Change password of an account
- Delete an account
- Verify JWT token

## Getting started

### Prerequisites

Go version 1.21.6 (latest version at the starting project moment).

### Installation

``` bash
git clone https://github.com/ltphat2204/goauth
cd goauth/
```

### Run
``` bash
docker compose up -d
```

### Configure
Let's take a look at the `docker-compose.yml` file. There are many adjustable things.

#### Database
- I use MySql as a database. You should change `ports` if there is any conflict. 

- In default, I use `root` user with the password `admin`. Any changes must match with the MySql account of the service below.

- Finally, I creating Database `AUTH_SERVICE` for saving data.Any changes must match with the MySql Database of the service below.

> For more setting, head over to [MySql Docker documentation](https://hub.docker.com/_/mysql)

#### Auth service
- In default, this service listens to port `5050`. You should change `ports` if there is any conflict. 

- `MYSQL_USER` always `root`, or another account you have setting above.

- `MYSQL_PASSWORD` and `MYSQL_DATABASE` must match `MYSQL_ROOT_PASSWORD` and `MYSQL_DATABASE` in the database setting.

- `MYSQL_URL` with the structure of `database:port` (where port is the MySql port setting above).

- `SECRET_KEY` is the key for signing JWT.

- `ADMIN_USERNAME` and `ADMIN_PASSWORD` is used for creating the super account of the service.

- `GIN_MODE` always `release` in production environment.


## Usage

### Getting JWT
You can call a GET method with Authorization header to the `/token` endpoint to get your token. In default, the token will expire in 30 minutes.

``` bash
curl --location 'http://localhost:5050/token' \
--header 'Authorization: Basic bHRwOjIyMDQ='
```

> Moreover, the text following `Basic ` is in format "`Username:Password`" in Base64 Encoding.

If the account exists and the password is matched, the response will be like:

``` json
{
    "success": true,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDgzNjY0NTksInVzZXJuYW1lIjoibHRwIn0.kKGCblGY19OIqRT-2MUpk_TEDBCv5XtPe285syhfTaU"
    }
}
```

### Verify JWT

You can call a GET method with Authorization header to the `/verify` endpoint to verify the token.

``` bash
curl --location 'http://localhost:5050/verify' \
--header 'Authorization: BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDcyNTE1NDcsInVzZXJuYW1lIjoibHRwIn0.UN3viM1PE0JIrOaNusRysdjS_KkL3eJa6bG889-yr8M'
```

> Note: The value of `Authorization` header must match the format "`BEARER token`", where `token` is the token to be verified.

If everything is OK, your response should be:

``` json
{
    "success": true,
    "data": "Valid token!"
}
```

### Get list of users

You can call a GET method to the `/` endpoint to get list of users.

**Require**: Only token of Admin user is accepted.

``` bash
curl --location 'http://localhost:5050/' \
--header 'Authorization: BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDcyNTE1NDcsInVzZXJuYW1lIjoibHRwIn0.UN3viM1PE0JIrOaNusRysdjS_KkL3eJa6bG889-yr8M'
```

Example of successful response:

``` json
{
    "success": true,
    "data": {
        "users": [
            {
                "username": "ltp",
                "created_at": "2024-02-06T12:58:31.298+07:00",
                "last_access_at": "2024-02-07T21:37:05.085+07:00"
            }
        ]
    }
}
```

### Create new user

You can call a POST method to `/` endpoint to create a new user.

**Require**: Only token of Admin user is accepted.

``` bash
curl --location 'http://localhost:5050/' \
--header 'Authorization: BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDcyNTE1NDcsInVzZXJuYW1lIjoibHRwIn0.UN3viM1PE0JIrOaNusRysdjS_KkL3eJa6bG889-yr8M' \
--header 'Content-Type: application/json' \
--data '{
    "username": "example",
    "password": "example"
}'
```

Example of successful response:

```json
{
    "success": true,
    "data": {
        "message": "User example created successfully!",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDgzNjY2NzgsInVzZXJuYW1lIjoiZXhhbXBsZSJ9.-cekYlhG1QTtNPJQ-c9_Uk8iBnVP0r8OIO0fqIgSQuA"
    }
}
```

> You can use that token instantly without calling the get token api.

### Get detail of user

You can get detail of a user by calling a GET method to `/:username` endpoint, where username is the username of the user you want to get detail.

**Require**: Only token of Admin user is accepted.

``` bash
curl --location 'http://localhost:5050/example' \
--header 'Authorization: BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDcyNTE1NDcsInVzZXJuYW1lIjoibHRwIn0.UN3viM1PE0JIrOaNusRysdjS_KkL3eJa6bG889-yr8M'
```

The successful response will be like:

```json
{
    "success": true,
    "data": {
            "username": "example",
            "created_at": "2024-02-07T21:46:22.373+07:00",
            "last_access_at": "2024-02-07T21:46:22.373+07:00"
        }
}
```

### Change password

You can call a PATCH method to `/:username` endpoint to change password of the username.

**Require**: The token must belong to the username which is being changed password.

``` bash
curl --location --request PATCH  'http://localhost:5050/ltp' \
--header 'Authorization: BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDcyNTE1NDcsInVzZXJuYW1lIjoibHRwIn0.UN3viM1PE0JIrOaNusRysdjS_KkL3eJa6bG889-yr8M'
--header 'Content-Type: application/json' \
--data '{
    "password": "another"
}'
```

The successful response will be like:

```json
{
    "success": true,
    "data": {
        "message": "Password update successfully!"
    }
}
```

### Delete a user

You can call a DELETE method to `/:username` endpoint to delete a user with the username.

``` bash
curl --location --request DELETE 'http://localhost:3030/example' \
--header 'Authorization: BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDcyNTE1NDcsInVzZXJuYW1lIjoibHRwIn0.UN3viM1PE0JIrOaNusRysdjS_KkL3eJa6bG889-yr8M'
```

The successful response will be like:

```json
{
    "success": true,
    "data": {
        "message": "User example deleted successfully!"
    }
}
```

> Note: You can not delete admin account

## Customization
Thanks to Clean Architect, it is pretty easy to customize the data model.

Head over to the `users/entity`, you can add more field to the `User` struct. If that field occurs in most situation, you can add it to `common/mysql_model.go`. Then, every model in `users/entity` will be added.

In term of working logic, take a look in the `users/business` and `users/transport` to modify the processing flow.

Finally, if you want to use another database, some changes in `common/mysql.go` and `users/storage` are needed.

Happy coding!