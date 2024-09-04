# Snipetbox for the book let's go book

This code shows all topics of all units of the book like:
Sending and receiving JSON
Working with SQL migrations
Managing background tasks
Performing partial updates and using optimistic locking
Permission-based authorization
Controlling CORS requests
Graceful shutdowns
Exposing application metrics
Automating build and deployment step

## Content table

- [Snipetbox for the book let's go book](#snipetbox-for-the-book-lets-go-book)
  - [Content table](#content-table)
  - [Installation](#installation)
    - [Use](#use)
    - [Project structure](#project-structure)
    - [Prerequisites](#prerequisites)
    - [Contribute](#contribute)

## Installation

You must install all dependencies with this command to install the code on your local machine.
```
go mod tidy
```
Before running the project, you must create a MySQL database with docker.
```
docker run --name MySQL -e MYSQL_ROOT_PASSWORD=@dmin1234 -e MYSQL_DATABASE=snippetbox -p 3306:3306 -d mysql:latest
```
Creating the tables and data examples in your database would be best.
```
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
You need to create certificates to run in HTTPS the project and create the path tls
```
cd tls
go run "/C/Program Files/Go/src/crypto/tls/generate_cert.go" --rsa-bits=2048 --host=localhost
```

### Use
Well, we are done installing everything. We must execute the following command to run the project.
```
go run ./cmd/web
```
You can send application parameters if you need to configure other parameters.
- addr: Http network address exapmple (-addr 127.0.0.1:8080)
- dsn: MySQL data source (-dsn user:pass@localhost:1234/snippetbox?parseTime=true)
- debug: To enable debug mode.

### Project structure

.
├── cmd
│   └── web
│       ├── context.go
│       ├── handlers.go
│       ├── helpers.go
│       ├── main.go        # Application entry point
│       ├── middleware.go
│       ├── routes.go
│       └── templates.go
├── internal
│   ├── assert
│   │   └── assert.go
│   ├── models
│   │   ├── errors.go
│   │   ├── snippets.go
│   │   └── users.go
│   └── validator
│       └── validator.go
├── tls
│   ├── cert.pem
│   └── key.pem
├── ui
│   ├── html
│   │   ├── pages
│   │   │  ├── about.gohtml
│   │   │  ├── account.gohtml
│   │   │  ├── create.gohtml
│   │   │  ├── home.gohtml
│   │   │  ├── login.gohtml
│   │   │  ├── password.gohtml
│   │   │  ├── sigup.gohtml
│   │   │  └── view.gohtml
│   │   ├── partials
│   │   │  └── nav.gohtml
│   │   └── base.gohtml
│   ├── static
│   │   ├── css
│   │   │  └── mains.css
│   │   ├── css
│   │   │  ├── favicon.ico
│   │   │  └── logo.png
│   │   └── js
│   │      └── mains.js
│   └── efs.go
├── go.mod
├── go.sum
└── README.md

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.23 o lastest)


### Contribute

- Fork the project
- Create a branch for your feature (git checkout -b feature/new-feature)
- Make your changes and commit (git commit -am 'Add new feature')
- Push your changes to your fork (git push origin feature/new-feature)
- Open a Pull Request
