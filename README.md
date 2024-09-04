# Snippetbox: Supporting Material for the "Let's Go" Book 📖
___
**This code shows all topics of all units of the book like**:
1. Sending and receiving JSON
1. Working with SQL migrations
1. Managing background tasks
1. Performing partial updates and using optimistic locking
1. Permission-based authorization
1. Controlling CORS requests
1. Graceful shutdowns
1. Exposing application metrics
1. Automating build and deployment step

___
## Table of Contents 📋

- [Snippetbox: Supporting Material for the "Let's Go" Book 📖](#snippetbox-supporting-material-for-the-lets-go-book-)
  - [Table of Contents 📋](#table-of-contents-)
  - [Installation 🛠️](#installation-️)
    - [Install dependencies](#install-dependencies)
    - [Install database](#install-database)
    - [Create tables](#create-tables)
    - [Create certificates](#create-certificates)
  - [Usage 🚀](#usage-)
  - [Project Structure 📂](#project-structure-)
  - [Prerequisites ✔️](#prerequisites-️)
  - [Contribute 🤝](#contribute-)

## Installation 🛠️

### Install dependencies
To install the code on your local machine, you need to install all the dependencies with the following command:
```go
go mod tidy
```

### Install database
Before running the project, you must create a MySQL database with Docker:
```bash
docker run --name MySQL -e MYSQL_ROOT_PASSWORD=@dmin1234 -e MYSQL_DATABASE=snippetbox -p 3306:3306 -d mysql:latest
```

### Create tables
It is recommended to create the tables and some sample data in your database:
```sql
-- Create a new UTF-8 `snippetbox` database.
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Create a new UTF-8 `test_snippetbox` database.
CREATE DATABASE test_snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Switch to using the `snippetbox` database.
USE snippetbox;

-- Create a `snippets` table.
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

-- Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets(created);

-- Add some dummy records (which we'll use in the next couple of chapters).
INSERT INTO snippets (title, content, created, expires) VALUES (
    'An old silent pond',
    'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'Over the wintry forest',
    'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'First autumn morning',
    'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);

CREATE TABLE sessions (
    token CHAR(43) PRIMARY KEY,
    data BLOB NOT NULL,
    expiry TIMESTAMP(6) NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);

CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);

```

### Create certificates
You need to create certificates to run the project in HTTPS and create the tls path:
``` bash
cd tls
go run "/C/Program Files/Go/src/crypto/tls/generate_cert.go" --rsa-bits=2048 --host=localhost
```

## Usage 🚀
Well, we are done installing everything. We must execute the following command to run the project.
```go
go run ./cmd/web
```
You can send application parameters if you need to configure other parameters.
- addr: Http network address exapmple (-addr 127.0.0.1:8080)
- dsn: MySQL data source (-dsn user:pass@localhost:1234/snippetbox?parseTime=true)
- debug: To enable debug mode.

## Project Structure 📂

```
.
├── cmd 📂
│   └── web 🕸️
│       ├── context.go 📄
│       ├── handlers.go 📄
│       ├── helpers.go 📄
│       ├── main.go 📄   🚀  (Application entry point)
│       ├── middleware.go 📄
│       ├── routes.go 📄
│       └── templates.go 📄
├── internal 📂
│   ├── assert ✅
│   │   └── assert.go 📄
│   ├── models 🗃️
│   │   ├── errors.go 📄
│   │   ├── snippets.go 📄
│   │   └── users.go 📄
│   └── validator ✔️
│       └── validator.go 📄
├── tls 🔒
│   ├── cert.pem 📄
│   └── key.pem 📄
├── ui 🖥️
│   ├── html 📄
│   │   ├── pages 📄
│   │   │   ├── about.gohtml 📄
│   │   │   ├── account.gohtml 📄
│   │   │   ├── create.gohtml 📄
│   │   │   ├── home.gohtml 📄
│   │   │   ├── login.gohtml 📄
│   │   │   ├── password.gohtml 📄
│   │   │   ├── signup.gohtml 📄
│   │   │   └── view.gohtml 📄
│   │   ├── partials 📄
│   │   │   └── nav.gohtml 📄
│   │   └── base.gohtml 📄
│   ├── static 📂
│   │   ├── css 🎨
│   │   │   └── main.css 📄
│   │   ├── img 🖼️
│   │   │   ├── favicon.ico 📄
│   │   │   └── logo.png 📄
│   │   └── js ✨
│   │       └── main.js 📄
│   └── efs.go 📄
├── go.mod 📄
├── go.sum 📄
└── README.md 📄
```

## Prerequisites ✔️

- [Go](https://golang.org/doc/install) (version 1.23 o lastest)


## Contribute 🤝

- Fork the project
- Create a branch for your feature (git checkout -b feature/new-feature)
- Make your changes and commit (git commit -am 'Add new feature')
- Push your changes to your fork (git push origin feature/new-feature)
- Open a Pull Request
