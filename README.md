# News Feed Server

In this project, I have created a server which listens for <br>
GET and POST requests at a particular port.<br>
<br>
User can get all news feeds<br>
And Post JSON of items to be added in News Feed's list <br>
<br>
To run this application

Install the dependaancies via
```
    go mod tidy
```

Run the application
```
    go run main.go
```

# To create from scratch

##   Chi Server in Go

Chi is a lightweight alternative to Gin
It provides simple wrappers around the standard libraries.

Install the dependancy:
```
    go get -u github.com/go-chi/chi/v5
```

Create a seperate file for 
<li> Domain </li>
<li>Chi Server </li>
<li>Http Handler Functions </li>

##  Domain
The domain/model contains necessary Add, Get and New methods

## Handlers
<li>These contain the basic Get and Post Handlers which are binded in chiserver.go</li>
<li>Get returns all the news items</li>
<li>Post adds a single news to list of news items</li>

## Chi Server
<li>This is the main file which sets up the Chi server for listening at port 3000
<li>NewServer() method returns the server and router object </li>
<li>setUpHandlers() method is used to bind get and Post methods to server </li>
<li>Start() method is used to make server listen at particular port</li>

<p>These endpoints have been tested on ThunderClient (VS Code extension) </p>

<p> GET: http://localhost:3000/newsfeed </p>

<p> POST: http://localhost:3000/newsfeed </p>
<p>    Body:  </p>
```json
{
    "title": "another title",
    "post": "another post"
}
```

## Zerolog and LumberJack

Install Dependancies: 
```
    go get -u github.com/rs/zerolog/log
    go get gopkg.in/natefinch/lumberjack.v2
```
To implement logging functionality in code, <br>
I used Zerolog and LumberJack.
Zerolog provides a structured way for writing logs.
Lumberjack helps writing logs to rolling files.
