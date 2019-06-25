package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	gql_company "github.com/Shahid4183/gql-company"
)

func main() {

	// using this flag user can specify which port to start this server on
	// by default it will be 8080 port
	port := flag.String("port", "8080", "Server port number")
	// using this flag user can specify whether to auto migrate database tables
	// by default it will be false
	autoMigrate := flag.Bool("auto-migrate", false, "auto migrate database schema")
	flag.Parse()

	// connect to database
	// by default it will try to connect to mysql server running on localhost:3306
	// using username = "root", password = "root" and database = company
	// make sure you have mysql running on port 3306 and it has company database
	// otherwise it will panic and code will not run
	if err := gql_company.ConnectToDatabase(); err != nil {
		fmt.Println("Error:", err)
		panic("Couldn't connect to database")
	}

	if *autoMigrate {
		gql_company.AutoMigrate()
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gql_company.NewExecutableSchema(gql_company.Config{Resolvers: &gql_company.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", *port)
	log.Fatal(http.ListenAndServe(":"+(*port), nil))
}
