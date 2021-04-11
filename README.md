# gimmy
# About

I was assigned to create a RESTful API involving a fictional characters game without using any library, framework or copy-paste and this is the result of a week-end work.
Altough it's not perfect I want to keep it like this to see my coding skills at that moment. For more info about the API requests and responses check the controllers.

# Installation instructions

* Clone the repository
* Run tests go test ./...
* Run the app go run main.go

# Notes

The server runs on port :8282. You can change the port inside main.go

# API 

POST /classes - creates a new classes

Request:
```
{"name": "Karate", "capacity": 26, "start_date": "2021-04-30", "end_date": "2021-05-01"}
```

Success response:
```
{"id":"458c0903-c394-4a63-8289-57cc5bf59305"}
```
