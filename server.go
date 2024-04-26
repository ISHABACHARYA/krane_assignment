package main

import (
	"eventManagemntSystem/graph"
	"eventManagemntSystem/postgres"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	// Connect to PostgreSQL database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("host")
	dbPort := os.Getenv("dbPort")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")

	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, dbPort, user, password, dbname)

	db := postgres.New(connStr)

	defer db.Close()
	// Read the SQL query from schema.sql file
	// sqlBytes, err := os.ReadFile("database.sql")
	// if err != nil {
	// 	log.Fatalf("Error reading schema.sql file: %v", err)
	// }
	// sqlQuery := string(sqlBytes)

	// // Execute the SQL query to create tables
	// _, err = db.Exec(sqlQuery)
	// if err != nil {
	// 	log.Fatalf("Error creating tables: %v", err)
	// }

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	userRepo := postgres.InitUserRepo(db)
	eventRepo := postgres.InitEventRepo(db)
	expenseRepo := postgres.InitExpenseRepo(db)
	r := &graph.Resolver{UserRepo: &userRepo, EventRepo: &eventRepo,ExpenseRepo: &expenseRepo}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: r}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
