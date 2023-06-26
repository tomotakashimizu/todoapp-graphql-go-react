package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/tomotakashimizu/todoapp-graphql-go-react/app/graph"
	"xorm.io/xorm"
)

const defaultPort = "8080"

func main() {
	connectionString := "user=postgres password=postgres dbname=testdb host=db port=5432 sslmode=disable"
	engine, err := xorm.NewEngine("postgres", connectionString)
	if err != nil {
		log.Fatalln("error - create engine: ", err)
	}

	err = engine.Ping()
	if err != nil {
		log.Fatalln("error - connect DB: ", err)
	}
	log.Println("success - connect DB")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: engine}}))

	// Add CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	handler := c.Handler(http.DefaultServeMux)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
