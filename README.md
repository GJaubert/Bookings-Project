# Bookings project - Go & back-end

> Gabriel García Jaubert  
> Bookings project - Go & back-end  
> 13/03/2023

This project is meant for learning Go and practice my back end skills.

It's a simple idea of a web page for booking different hotel rooms. The idea of the web is not the important part, the real purpose is to practice things like CSRF token, sessions, routing, repository pattern, security, testing, etc.

To run the project, go the root of the project and type:

```
$ go run cmd/web/main.go cmd/web/routes.go cmd/web/middleware.go
```

To run the tests:

```$ cd {directory you want to test}```

```$ go test```

The page should be accesible from port :8080

Transactions with the database are not possible because the db has no public endpoint.

The frontend of this page was made with bootstrap code examples.