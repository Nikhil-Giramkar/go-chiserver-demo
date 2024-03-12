#   Chi Server in Go

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
