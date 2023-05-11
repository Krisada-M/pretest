package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EmployeeAccount is main model
type EmployeeAccount struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	EmployeeID     string             `json:"employee_id" bson:"employee_id,omitempty"`
	Firstname      string             `json:"firstname" bson:"firstname,omitempty" validate:"required,max=50"`
	Lastname       string             `json:"lastname" bson:"lastname,omitempty" validate:"required,max=50"`
	Position       string             `json:"position" bson:"position,omitempty" validate:"required,max=50"`
	Salary         uint               `json:"salary" bson:"salary,omitempty" validate:"required"`
	Yearexperience string             `json:"year_experience" bson:"year_experience,omitempty" validate:"required,max=50"`
	StartDate      string             `json:"start_date" bson:"start_date,omitempty" validate:"required,max=50"`
	CreateAt       time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt       time.Time          `json:"update_at,omitempty" bson:"update_at,omitempty"`
}
