package models

import (
	"time"

	"github.com/Subodhsanjayband/event_manager/db"
)

type Event struct {
	ID          int64
	Name        string    `binding: "required"`
	Description string    `binding: "required"`
	Loaction    string    `binding: "required"`
	DateTime    time.Time `binding: "required"`
	UserID      int64
}

var events = []Event{}

func (e *Event) Save() error {
	insertQuery := `
	INSERT INTO events (name,description,location,dateTime,user_id)
	 values(?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(insertQuery)

	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Loaction, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId() /// check this point not returning correct id

	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	selectQuery := `SELECT * FROM events`
	rows, err := db.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Loaction, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	selectQuery := `SELECT * FROM events WHERE id = ?`

	row := db.DB.QueryRow(selectQuery, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Loaction, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil

}

func (e Event) Update() (*Event, error) {
	updateQuery := `UPDATE events SET name = ?, description = ?, location=?,dateTime=? WHERE id = ?`
	stmt, err := db.DB.Prepare(updateQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Loaction, e.DateTime, e.ID)
	if err != nil {
		return nil, err
	}
	return GetEventByID(e.ID)

}
func (e Event) Delete() error {
	deleteQuery := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(deleteQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	if err != nil {
		return err
	}
	return nil
}
