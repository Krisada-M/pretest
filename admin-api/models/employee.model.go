package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EmployeeAccount is main model
type EmployeeAccount struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	EmployeeID     string             `json:"employee_id"`
	Firstname      string             `json:"firstname"`
	Lastname       string             `json:"lastname"`
	Position       string             `json:"position"`
	Salary         string             `json:"salary"`
	Yearexperience string             `json:"year_experience"`
	StartDate      string             `json:"start_date"`
	CreateAt       time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt       time.Time          `json:"update_at,omitempty" bson:"update_at,omitempty"`
}
