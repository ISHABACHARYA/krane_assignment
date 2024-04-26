package postgres

import (
	"database/sql"
	"eventManagemntSystem/model"
	"fmt"
	"log"

	"github.com/doug-martin/goqu/v9"
)

type ExpenseRepo struct {
	DB *sql.DB
}

func InitExpenseRepo(db *sql.DB) ExpenseRepo {
	return ExpenseRepo{DB: db}
}

func (u *ExpenseRepo) CreateExpense(userID string, expense model.ExpenseInput) (model.Expense, error) {
	gq := GoquNew(u.DB)
	sql, args, _ := gq.Insert("Expense").Rows(expense).Returning("id").ToSQL()
	fmt.Print(sql)
	result, err := u.DB.Exec(sql, args...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result from , %+v", result)
	var eventID int64
	ferr := u.DB.QueryRow(sql, args...).Scan(&eventID)
	if ferr != nil {
		return model.Expense{}, err
	}

	// Retrieve the inserted user from the database using the ID
	var newEvent model.Expense
	query := gq.From("Expense").Select("id","eventid","type","amount","name","description").Where(goqu.Ex{"id": eventID})
	sql, args, _ = query.ToSQL()
	err = u.DB.QueryRow(sql, args...).
		Scan(&newEvent.ID, &newEvent.EventID, &newEvent.Type, &newEvent.Amount, &newEvent.Name, &newEvent.Description)
	if err != nil {
		return model.Expense{}, err
	}
	return newEvent, nil
}

func (u *ExpenseRepo) GetExpensesByEventId(eventId string) (*[]model.Expense, error) {
    gq := GoquNew(u.DB)
    var expenses []model.Expense
    query := gq.From("Expense").Select("id", "eventid", "type", "amount", "name", "description").Where(goqu.Ex{"eventid": eventId})
    sql, args, _ := query.ToSQL()
    rows, err := u.DB.Query(sql, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        var newExpense model.Expense
        err := rows.Scan(&newExpense.ID, &newExpense.Name, &newExpense.EventID, &newExpense.Type, &newExpense.Amount, &newExpense.Description)
        if err != nil {
            return nil, err
        }
        expenses = append(expenses, newExpense)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return &expenses, nil
}
