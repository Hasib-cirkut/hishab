## Stack

- golang
- postgresql
- gorm

## Setup and run

### Database

The DB is ran using a docker container.

1. Create a `.env` file in the same directory as go.mod
2. Copy over the file content of `.env.example` to `.env`
3. Change .env file contents according to your choice
4. Run the following command

```bash
docker-compose up
```

### API

The api is a small golang app that uses gin framework.

1. Run the following command

```bash
go run cmd/hishab/main.go
```

The first time, the app will setup the DB with predefined tables and categories and other stuffs.
