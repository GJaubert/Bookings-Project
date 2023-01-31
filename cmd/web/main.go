package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gjaubert/bookings-project/pkg/config"
	"github.com/gjaubert/bookings-project/pkg/handlers"
	"github.com/gjaubert/bookings-project/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const PORT_NUMBER = ":8080"

var session *scs.SessionManager
var app config.AppConfig

func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println("Starting app in port 8080")
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	// _ = http.ListenAndServe(":8080", nil)
	srv := &http.Server{
		Addr:    PORT_NUMBER,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

/*
func Divide(w http.ResponseWriter, r *http.Request) {
	result, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}
	_, _ = fmt.Fprintf(w, "100 divided by 0 is %f\n", result)
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}
*/
