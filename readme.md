# Auth
An Account Creation, Account Deletion, User Authentication, Authorizatin And Session Handling API

## How TO Run

### Required
- Postgre
- Redis

### Env configuration
you should open the .env file and modify it

#### Run
clone the repo

```
git clone https://github.com/ftsog/ecom
```

move into the program directory

```
cd ecom 
```

run the program

```
$ go run .
```

NOTE: make sure you have the postgres server, redis server and the .env all setup before running the program

## API Techstack
- Go
- Postgres
- Redis

## API Endpoints

- /api/auth/register
- /api/auth/verify/{username}/{token}
- /api/auth/login
- /api/auth/logout
- /api/auth/delete
- /api/home

## Usage
* Register Functon API Callback(POST REQUEST)
- 127.0.0.1:8080/api/auth/register

```
{
    "username": "example12345",
    "email": "example12345@example.com",
    "password": "Example12345@"
}
```

Response:

```
{
    "Status": 200,
    "Message": "Accounted Successfully Created, Account verification link sent to your email",
    "Path": "/api/auth/register",
    "Redirect": "/api/login"
}
```

* Login Functon API Callback (POST REQUEST)
- 127.0.0.1:8080/api/auth/login

```
{
    "username": "example12345",
    "password": "Example12345@"
}
```

Response:

```
{
    "Status": 200,
    "Message": "Succesfully Login",
    "Path": "/api/auth/login",
    "Redirect": "/api/home"
}
```

* With the above giving example on how to use the API and by looking into the program you should be able to use the rest of the API
