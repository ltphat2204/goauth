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

- `SECRET_KEY` is the key for signing JWT

- `ADMIN_USERNAME` and `ADMIN_PASSWORD` is used for creating the super account of the service.

- `GIN_MODE` always `release` in production environment.