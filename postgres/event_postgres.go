package postgres

import (
	"database/sql"
	"eventManagemntSystem/model"
	"fmt"
	"log"

	"github.com/doug-martin/goqu/v9"
)

type EventRepo struct {
	DB *sql.DB
}

func InitEventRepo(db *sql.DB) EventRepo {
	return EventRepo{DB: db}
}

func (u *EventRepo) CreateEvent(userID string, event model.EventInput) (model.Event, error) {
	gq := GoquNew(u.DB)
	sql, args, _ := gq.Insert("Event").Rows(event).Returning("id").ToSQL()
	fmt.Print(sql)
	result, err := u.DB.Exec(sql, args...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result from , %+v", result)
	var eventID int64
	ferr := u.DB.QueryRow(sql, args...).Scan(&eventID)
	if ferr != nil {
		return model.Event{}, err
	}

	fmt.Printf("eventId >>>> %+v", eventID)
	// Retrieve the inserted user from the database using the ID
	var newEvent model.Event
	query := gq.From("Event").Select("id", "name", "startdate", "enddate", "createdat", "updatedat").Where(goqu.Ex{"id": eventID})
	sql, args, _ = query.ToSQL()
	err = u.DB.QueryRow(sql, args...).
		Scan(&newEvent.ID, &newEvent.Name, &newEvent.StartDate, &newEvent.EndDate, &newEvent.CreatedAt, &newEvent.UpdatedAt)
	if err != nil {
		return model.Event{}, err
	}
	fmt.Print("\n Event Added success. New Event: ", newEvent)
	// Insert into userEvent table
	userEvent := model.UserEvent{
		UserID:  &userID,
		EventID: &newEvent.ID,
		Role:    model.UserRoleAdmin,
	}

	{
		sql, args, _ := gq.Insert(("UserEvent")).Rows(userEvent).ToSQL()
		fmt.Println(sql)
		_, err := u.DB.Exec(sql, args...)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Print("\nInsert Result is", result)

	return newEvent, nil
}

func (u *EventRepo) GetEventById(id string) (*model.Event, error) {
	gq := GoquNew(u.DB)
	var newEvent model.Event
	query := gq.From("Event").Select("id", "name", "startdate", "enddate", "createdat", "updatedat").Where(goqu.Ex{"id": id})
	sql, args, _ := query.ToSQL()
	err := u.DB.QueryRow(sql, args...).
		Scan(&newEvent.ID, &newEvent.Name, &newEvent.StartDate, &newEvent.EndDate, &newEvent.CreatedAt, &newEvent.UpdatedAt)
	if err != nil {
		return &model.Event{}, err
	}
	return &newEvent, nil
}
