# gimmy

# About

I was assigned to create a RESTful API involving a fictional characters game without using any library, framework or
copy-paste and this is the result of a week-end work. Altough it's not perfect I want to keep it like this to see my
coding skills at that moment. For more info about the API requests and responses check the controllers.

# Installation instructions

* Clone the repository
* Run tests go test ./...
* Run the app go run main.go

# Notes

The server runs on port :8282. You can change the port inside main.go

# API

**POST** /classes - creates a new class

Sample request:

```
{"name": "Karate", "capacity": 26, "start_date": "2021-04-30", "end_date": "2021-05-01"}
```

Sample response:

Returns the ID of the created class object

```
{"id":"458c0903-c394-4a63-8289-57cc5bf59305"}
```

**GET** /classes

Sample response:

```
{"classes":[{"id":"458c0903-c394-4a63-8289-57cc5bf59305","name":"Karate","start_date":"2021-04-30","end_date":"2021-05-01","capacity":26},{"id":"53ee2b12-2fab-4f32-86f7-183821096da5","name":"Salsa","start_date":"2021-03-21","end_date":"2021-05-01","capacity":30}]}
```

**POST** /bookings - creates a new booking

Sample request:

```
{"name": "Andrei", "date": "2021-05-01", "class_id": "458c0903-c394-4a63-8289-57cc5bf59305"}
```

Sample response:

Returns the ID of the created booking object.

```
{"id":"b719bda8-86d1-4929-81c7-cd82369e037d"}
```