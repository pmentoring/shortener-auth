# Golang links shortener

## Run the app first time

You can run the app with the command following:
```bash
make init
```
That will compile & runs web server and proceed migrations.

To stop app use:
```bash
make down
```
To get into container terminal:
```bash
make exec
```

## Migrations
Run the migrations with command:

```bash
make migrate
```

Create new migration with command:
```bash
make migration
```

## Auth
register - get password and login (must be unique), create user, and return jwt with expires_at
login - get password and login, return jwt with expires_at
jwt renew - get jwt and create new one (difference is created_at field)
jwt:
    fields:
        userId: int
        issued_at: datetime
        expires_at: datetime
    lifetime: 3h