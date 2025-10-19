# GOLANG MYSQL CRUD API

Sample crud operation using Golang and MySql.

## Prerequisites

Docker, Git, Go. Check `conf-dev.yml` from a configuration example.


### Installing

clone the repo, then fetch dependencies and u're good to go.

```
    git clone git@github.com/golang-ddd-postgreql-auth.git $GOPATH/src/github.com/golang-ddd-postgreql-auth
    cd $GOPATH/src/github.com/golang-ddd-postgreql-auth
    export GO111MODULE=on;
    go mod vendor;
    go mod download;
    go mod tidy
    make dev
```

### Build locally

```
    git clone git@github.com/golang-ddd-postgreql-auth.git $GOPATH/src/github.com/golang-ddd-postgreql-auth
    cd $GOPATH/src/github.com/golang-ddd-postgreql-auth
    export GO111MODULE=on;
    go mod vendor;
    go mod download;
    go mod tidy
    make build
```


### Running the tests

```
    export GO111MODULE=on;
    go mod vendor;
    go mod download;
    go mod tidy
    make test
```

### Usage

## API ENDPOINTS

### All Posts
- Path : `/posts`
- Method: `GET`
- Response: `200`

### Create Post
- Path : `/posts`
- Method: `POST`
- Fields: `title, content`
- Response: `201`

### Details a Post
- Path : `/posts/{id}`
- Method: `GET`
- Response: `200`

### Update Post
- Path : `/posts/{id}`
- Method: `PUT`
- Fields: `title, content`
- Response: `200`

### Delete Post
- Path : `/posts/{id}`
- Method: `DELETE`
- Response: `204`

## Required Packages
- Database
    * [MySql](https://github.com/go-sql-driver/mysql)
- Routing
    * [chi](https://github.com/go-chi/chi)



### Notes on Architecture

![Sequencial Diagram](doc/diagram.png)