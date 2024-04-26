package postgres

import (
	"database/sql"
	"eventManagemntSystem/model"
	"fmt"
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type UsersRepo struct {
	DB *sql.DB
}

func InitUserRepo(db *sql.DB) UsersRepo {
	return UsersRepo{DB: db}
}

var tableName string = "User"

func (u *UsersRepo) GetUserByID(id string) (*model.User, error) {
	var user model.User
	gq := GoquNew(u.DB)
	sql, args, _ := gq.Select("*").From(tableName).Where(goqu.Ex{"id": id}).ToSQL()
	rows, err := u.DB.Query(sql, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.PhoneNumber, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("user with id %s not found", id)
	}

	return &user, nil
}

func (u *UsersRepo) CreateUser(user model.UserInput) (model.User, error) {
	var newUser model.User
	fmt.Print(tableName, user)
	gq := GoquNew(u.DB)
	sql, args, _ := gq.Insert(tableName).Rows(user).Returning("id").ToSQL()
	fmt.Print(sql)
	result, err := u.DB.Exec(sql, args...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\nInsert Result is", result)

	var userID int64
	ferr := u.DB.QueryRow(sql, args...).Scan(&userID)
	if ferr != nil {
		return model.User{}, err
	}

	fmt.Println("User ID>>>>>", userID)
	// Retrieve the inserted user from the database using the ID
	query := gq.From(tableName).Where(goqu.Ex{"id": userID})
	sql, args, _ = query.ToSQL()
	err = u.DB.QueryRow(sql, args...).
		Scan(&newUser.ID, &newUser.Username, &newUser.PhoneNumber, &newUser.Email, &newUser.CreatedAt, &newUser.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}

	return newUser, nil
}

func (u *UsersRepo) UpdateUser(id *string, user model.UserInput) (model.User, error) {
	var updatedUser model.User
	gq := GoquNew(u.DB)

	// Generate the SQL query for updating the user
	sql, args, _ := gq.Update(tableName).
		Set(goqu.Record{
			"username":    user.Username,
			"email":       user.Email,
			"phonenumber": user.PhoneNumber,
			"updatedat":   time.Now(), // Assuming you want to update the 'updatedat' field
		}).
		Where(goqu.Ex{"id": id}).
		ToSQL()

	// Execute the update query
	_, err := u.DB.Exec(sql, args...)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the updated user from the database using the provided ID
	query := gq.From(tableName).Where(goqu.Ex{"id": id})
	sql, args, _ = query.ToSQL()
	err = u.DB.QueryRow(sql, args...).
		Scan(&updatedUser.ID, &updatedUser.Username, &updatedUser.PhoneNumber, &updatedUser.Email, &updatedUser.CreatedAt, &updatedUser.UpdatedAt)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to retrieve updated user: %v", err)
	}

	return updatedUser, nil
}
