package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/ofonimefrancis/problemsApp/config"
	"github.com/ofonimefrancis/problemsApp/features/comments"
	"github.com/ofonimefrancis/problemsApp/features/likes"
	"github.com/ofonimefrancis/problemsApp/features/problems"
	"github.com/ofonimefrancis/problemsApp/features/users"
)

func main() {
	config.Init()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://ibisubizo.com/*", "*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)

	r.Mount("/api/auth", users.AuthRoutes())
	r.Mount("/api/problems", problems.Routes())
	r.Mount("/api/comments", comments.Routes())
	r.Mount("/api/users", users.UserRoutes())
	r.Mount("/api/likes", likes.Routes())

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}

	log.Printf("Now starting server on port %s", config.Get().Port)

	PORT := fmt.Sprintf(":%s", config.Get().Port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Stopping server...")
		os.Exit(1)
	}()
	log.Println(http.ListenAndServe(PORT, r))
}
