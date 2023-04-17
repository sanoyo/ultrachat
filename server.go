package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sanoyo/ultrachat/aws"
	"github.com/sanoyo/ultrachat/graph"
	"github.com/sanoyo/ultrachat/repository"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sql.Open("mysql", "wuser:password@tcp(localhost:3307)/ultrachat?parseTime=true")
	if err != nil {
		fmt.Printf("failed to connect to database: %v\n", err)
		return
	}
	defer db.Close()

	dynamoClient := aws.NewDynamoClient("http://localhost:8000")
	spaces := repository.NewSpaceRepository(db)
	users := repository.NewUserRepository(db)

	resolver := graph.NewResolver(dynamoClient, *spaces, *users)
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
