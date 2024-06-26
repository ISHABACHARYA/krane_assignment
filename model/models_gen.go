// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Event struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type EventInput struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type Expense struct {
	ID          string      `json:"id"`
	EventID     string      `json:"eventId"`
	Type        ExpenseType `json:"type"`
	Amount      int         `json:"amount"`
	Name        string      `json:"name"`
	Description *string     `json:"description,omitempty"`
}

type ExpenseInput struct {
	EventID     string      `json:"eventId"`
	Type        ExpenseType `json:"type"`
	Amount      int         `json:"amount"`
	Name        string      `json:"name"`
	Description *string     `json:"description,omitempty"`
}

type Mutation struct {
}

type Query struct {
}

type Session struct {
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	Status     *string   `json:"status,omitempty"`
	Visibility *bool     `json:"visibility,omitempty"`
	Event      string    `json:"event"`
}

type User struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber *string   `json:"phoneNumber,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UserEvent struct {
	ID      string   `json:"id"`
	UserID  *string  `json:"userId,omitempty"`
	EventID *string  `json:"eventId,omitempty"`
	Role    UserRole `json:"role"`
}

type UserEventInput struct {
	UserID  string   `json:"userID"`
	EventID string   `json:"eventID"`
	Role    UserRole `json:"role"`
}

type UserInput struct {
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

type ExpenseType string

const (
	ExpenseTypeVenue       ExpenseType = "VENUE"
	ExpenseTypeCatering    ExpenseType = "CATERING"
	ExpenseTypeDecorations ExpenseType = "DECORATIONS"
	ExpenseTypeMisc        ExpenseType = "MISC"
	ExpenseTypePetty       ExpenseType = "PETTY"
)

var AllExpenseType = []ExpenseType{
	ExpenseTypeVenue,
	ExpenseTypeCatering,
	ExpenseTypeDecorations,
	ExpenseTypeMisc,
	ExpenseTypePetty,
}

func (e ExpenseType) IsValid() bool {
	switch e {
	case ExpenseTypeVenue, ExpenseTypeCatering, ExpenseTypeDecorations, ExpenseTypeMisc, ExpenseTypePetty:
		return true
	}
	return false
}

func (e ExpenseType) String() string {
	return string(e)
}

func (e *ExpenseType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ExpenseType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ExpenseType", str)
	}
	return nil
}

func (e ExpenseType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserRole string

const (
	UserRoleAdmin    UserRole = "ADMIN"
	UserRoleManager  UserRole = "MANAGER"
	UserRoleAttendee UserRole = "ATTENDEE"
)

var AllUserRole = []UserRole{
	UserRoleAdmin,
	UserRoleManager,
	UserRoleAttendee,
}

func (e UserRole) IsValid() bool {
	switch e {
	case UserRoleAdmin, UserRoleManager, UserRoleAttendee:
		return true
	}
	return false
}

func (e UserRole) String() string {
	return string(e)
}

func (e *UserRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserRole", str)
	}
	return nil
}

func (e UserRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
