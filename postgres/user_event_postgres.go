package postgres

import (
	"eventManagemntSystem/model"
	"fmt"
	"log"

	"github.com/doug-martin/goqu/v9"
)

func (u *EventRepo) AddUserToEvent(userID string, userEvent model.UserEventInput) (*model.UserEvent, error) {
	gq := GoquNew(u.DB)

	sql, args, _ := gq.Insert(("UserEvent")).Rows(userEvent).Returning("id").ToSQL()
	fmt.Println(sql)
	_, err := u.DB.Exec(sql, args...)
	if err != nil {
		log.Fatal(err)
	}
	var newEvent model.UserEvent
	var userEventID int64
	ferr := u.DB.QueryRow(sql, args...).Scan(&userEventID)
	if ferr != nil {
		return &model.UserEvent{}, err
	}

	query := gq.From("UserEvent").Where(goqu.Ex{"id": userEventID})
	sql, args, _ = query.ToSQL()
	err = u.DB.QueryRow(sql, args...).
		Scan(&newEvent.ID, &newEvent.UserID, &newEvent.EventID, &newEvent.Role)

	if err != nil {
		return &model.UserEvent{}, err
	}

	return &newEvent, nil
}

func (u *EventRepo) UpdateUserEvent(userID string, userEvent model.UserEventInput) (*model.UserEvent, error) {
	gq := GoquNew(u.DB)

	sql, args, _ := gq.Update(("UserEvent")).Where(goqu.Ex{"userid":userEvent.UserID,"eventId":userEvent.EventID}).ToSQL()
	fmt.Println(sql)
	_, err := u.DB.Exec(sql, args...)
	if err != nil {
		log.Fatal(err)
	}
	var newEvent model.UserEvent
	var userEventID int64
	ferr := u.DB.QueryRow(sql, args...).Scan(&userEventID)
	if ferr != nil {
		return &model.UserEvent{}, err
	}

	query := gq.From("UserEvent").Where(goqu.Ex{"id": userEventID})
	sql, args, _ = query.ToSQL()
	err = u.DB.QueryRow(sql, args...).
		Scan(&newEvent.ID, &newEvent.UserID, &newEvent.EventID, &newEvent.Role)

	if err != nil {
		return &model.UserEvent{}, err
	}

	return &newEvent, nil
}
func (u *EventRepo) GetUserEventByUserAndEventID(userID string, eventID string) (model.UserEvent, error) {
	gq := GoquNew(u.DB)
	fmt.Println("from get user event by user and event id")

	sql, args, _ := gq.Select("*").From("UserEvent").Where(goqu.Ex{"userid": userID,"eventid": eventID}).ToSQL()
	fmt.Println(sql)
	rows, err := u.DB.Query(sql, args...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Rows ",rows)
	defer rows.Close()
	var userEvent model.UserEvent
	if rows.Next() {
		err := rows.Scan(&userEvent.ID, &userEvent.UserID,&userEvent.EventID,&userEvent.Role)
		if err != nil {
			return model.UserEvent{}, err
		}
		} else {

		return model.UserEvent{}, fmt.Errorf("UserEvent with user id %s and eventId: %s not found", userID,eventID)
	}
	return userEvent, nil
}
