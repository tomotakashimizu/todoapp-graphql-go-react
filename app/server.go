package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"github.com/tomotakashimizu/todoapp-graphql-go-react/app/graph"
	"xorm.io/xorm"
)

const defaultPort = "8080"

func main() {
	connectionString := "user=test_user password=testtest dbname=test_db host=localhost port=5432 sslmode=disable"
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

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
