package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"eventManagemntSystem/model"
	"eventManagemntSystem/postgres"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	User      *model.User
	Event     *[]model.Event
	UserRepo  *postgres.UsersRepo
	EventRepo *postgres.EventRepo
	ExpenseRepo *postgres.ExpenseRepo
}
