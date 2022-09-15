package main

import (
	"github.com/ezrx/notification-service/graph/db"
	"github.com/ezrx/notification-service/graph/firebase"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ezrx/notification-service/graph"
	"github.com/ezrx/notification-service/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	resolver := &graph.Resolver{
		NotificationRepo:       db.NewNotificationRepo(),
		FirebaseRepo:           db.NewFirebaseRepo(),
		NotificationSender:     firebase.NewNotificationSender(),
		SimpleNotificationRepo: db.NewSimpleNotificationRepo(),
		OrderRepo:              db.NewOrderRepo(),
	}
	router := chi.NewRouter()
	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{
			"*",
		},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
