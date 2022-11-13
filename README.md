# go-gin-mysql
Go : MySQL database and driver usage for gin based RESTful APIs.

This repository contains developing of RESTful web service APIs in Go Programming Language using Gin Web Framework (Gin) and uses generic Go SQL Database, MySQL Driver packages for DB CRUD operations.

**Prerequisite:**
- Make sure GO already installed and had working environment.<br>
while developing this code go version used,<br>
```
$go version
go version go1.19.3 windows/amd64 
```
- Male sure MySQL installed and tested with `MySQL Workbench` or `MySQL Command Line Client`<br>
while developing this code MySQL version used,<br>
```
Server version: 8.0.31 MySQL Community Server - GPL
```

**Steps to Run:**
- Run `go run main.go` in terminal
- Go to another command terminal and test the REST APIs using CURL (or use REST client like Postman).

**Test Results:**<br>
CRUD REST APIs test using CURL and it's Results:
```
$curl -X GET http://localhost:8080/ping
"Hello.. Welcome!"
$curl -X GET http://localhost:8080/books
{"books":null}
$curl -X GET http://localhost:8080/books | python -m json.tool
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    14  100    14    0     0   1004      0 --:--:-- --:--:-- --:--:--  1076
{
    "books": null
}

$curl -X POST http://localhost:8080/books  -H "Content-Type: application/json" -d @book1.json
{"id":"6dcd7ac3-3fbb-4a06-969c-9b25f711a1b0","name":"C Programming Language","author":"Brian W. Kernighan, Dennis M. Ritchie","price":57.09,"pages":272,"date_published":"1988-03-22"}
$curl -X GET http://localhost:8080/books | python -m json.tool
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   194  100   194    0     0  12109      0 --:--:-- --:--:-- --:--:-- 12933
{
    "books": [
        {
            "id": "6dcd7ac3-3fbb-4a06-969c-9b25f711a1b0",
            "name": "C Programming Language",
            "author": "Brian W. Kernighan, Dennis M. Ritchie",
            "price": 57.09,
            "pages": 272,
            "date_published": "1988-03-22"
        }
    ]
}

$curl -X POST http://localhost:8080/books  -H "Content-Type: application/json" -d @book2.json
{"id":"4b70a1cd-57fc-4a97-93f5-0bf57ef1994d","name":"The C++ Programming Language","author":"Bjarne Stroustrup","price":64.31,"pages":1376,"date_published":"2013-05-09"}
$curl -X GET http://localhost:8080/books | python -m json.tool
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   364  100   364    0     0  21328      0 --:--:-- --:--:-- --:--:-- 22750
{
    "books": [
        {
            "id": "4b70a1cd-57fc-4a97-93f5-0bf57ef1994d",
            "name": "The C++ Programming Language",
            "author": "Bjarne Stroustrup",
            "price": 64.31,
            "pages": 1376,
            "date_published": "2013-05-09"
        },
        {
            "id": "6dcd7ac3-3fbb-4a06-969c-9b25f711a1b0",
            "name": "C Programming Language",
            "author": "Brian W. Kernighan, Dennis M. Ritchie",
            "price": 57.09,
            "pages": 272,
            "date_published": "1988-03-22"
        }
    ]
}

$curl -X PATCH http://localhost:8080/books/4b70a1cd-57fc-4a97-93f5-0bf57ef1994d -H "Content-Type: application/json" -d @update2.json
{"id":"4b70a1cd-57fc-4a97-93f5-0bf57ef1994d","name":"The C++ Programming Language - 2nd Edition","author":"B Stroustrup ","price":65.01,"pages":1399,"date_published":"2017-07-07"}
$curl -X GET http://localhost:8080/books | python -m json.tool
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   374  100   374    0     0  22446      0 --:--:-- --:--:-- --:--:-- 24933
{
    "books": [
        {
            "id": "4b70a1cd-57fc-4a97-93f5-0bf57ef1994d",
            "name": "The C++ Programming Language - 2nd Edition",
            "author": "B Stroustrup ",
            "price": 65.01,
            "pages": 1399,
            "date_published": "2017-07-07"
        },
        {
            "id": "6dcd7ac3-3fbb-4a06-969c-9b25f711a1b0",
            "name": "C Programming Language",
            "author": "Brian W. Kernighan, Dennis M. Ritchie",
            "price": 57.09,
            "pages": 272,
            "date_published": "1988-03-22"
        }
    ]
}

$curl -X DELETE http://localhost:8080/books/4b70a1cd-57fc-4a97-93f5-0bf57ef1994d

$curl -X GET http://localhost:8080/books | python -m json.tool
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   194  100   194    0     0  11701      0 --:--:-- --:--:-- --:--:-- 12933
{
    "books": [
        {
            "id": "6dcd7ac3-3fbb-4a06-969c-9b25f711a1b0",
            "name": "C Programming Language",
            "author": "Brian W. Kernighan, Dennis M. Ritchie",
            "price": 57.09,
            "pages": 272,
            "date_published": "1988-03-22"
        }
    ]
}
```
