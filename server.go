package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sanoyo/ultrachat/aws"
	"github.com/sanoyo/ultrachat/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dynamoClient := aws.NewDynamoClient("http://localhost:8000")
	resolver := graph.NewResolver(dynamoClient)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{
			Resolvers: resolver,
		},
	))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
