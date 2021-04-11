# gimmy

# About

Gimmy is simple API to manage creating classes and bookings for a fictional gym written in Golang

# Installation instructions

You should have at least Go 1.13 installed since this project uses Go modules

```shell
git clone git@github.com:zuzuleinen/gimmy.git
go run main.go
```

# Notes

The server runs on port :8282. You can change the port inside main.go

# API

**POST** /classes - creates a new class

Sample request:

```json
{
  "name": "Karate",
  "capacity": 26,
  "start_date": "2021-04-30",
  "end_date": "2021-05-01"
}
```

Sample response:

Returns the ID of the created class object

```json
{
  "id": "458c0903-c394-4a63-8289-57cc5bf59305"
}
```

**GET** /classes

Sample response:

```json
{
  "classes": [
    {
      "id": "458c0903-c394-4a63-8289-57cc5bf59305",
      "name": "Karate",
      "start_date": "2021-04-30",
      "end_date": "2021-05-01",
      "capacity": 26
    },
    {
      "id": "53ee2b12-2fab-4f32-86f7-183821096da5",
      "name": "Salsa",
      "start_date": "2021-03-21",
      "end_date": "2021-05-01",
      "capacity": 30
    }
  ]
}
```

**POST** /bookings - creates a new booking

Sample request:

```json
{
  "name": "Andrei",
  "date": "2021-05-01",
  "class_id": "458c0903-c394-4a63-8289-57cc5bf59305"
}
```

Sample response:

Returns the ID of the created booking object.

```json
{
  "id": "b719bda8-86d1-4929-81c7-cd82369e037d"
}
```
# Error response

Usually error responses return an 4XX code with following format:

```json
{"message":"`end_date` is required"}
```

