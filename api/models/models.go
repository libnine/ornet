package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName   string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName    string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Role        string             `json:"role,omitempty" bson:"role,omitempty"`
	Company     string             `json:"company,omitempty" bson:"company,omitempty"`
	Phone       string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	Username    string             `json:"username,omitempty" bson:"username,omitempty"`
	DateCreated *time.Time         `json:"date_created,omitempty" bson:"date_created,omitempty"`
}

type Eq struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DateCreated *time.Time         `json:"date_created,omitempty" bson:"date_created,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	ImageURL    string             `json:"imageurl,omitempty" bson:"imageurl,omitempty"`
}

type Case struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CaseID       string             `json:"case_id,omitempty" bson:"case_id,omitempty"`
	CaseDate     *time.Time         `json:"case_date,omitempty" bson:"case_date,omitempty"`
	BodyPart     string             `json:"body_part,omitempty" bson:"body_part,omitempty"`
	Description  string             `json:"description,omitempty" bson:"description,omitempty"`
	LeftRight    string             `json:"left_right,omitempty" bson:"left_right,omitempty"`
	Hospital     string             `json:"hospital,omitempty" bson:"hospital,omitempty"`
	Doctor       string             `json:"doctor,omitempty" bson:"doctor,omitempty"`
	Patient      string             `json:"patient,omitempty" bson:"patient,omitempty"`
	PatientDob   string             `json:"patient_dob,omitempty" bson:"patient_dob,omitempty"`
	Rep          string             `json:"rep,omitempty" bson:"rep,omitempty"`
	Subspecialty string             `json:"subspecialty,omitempty" bson:"subspecialty,omitempty"`
	DateCreated  *time.Time         `json:"date_created,omitempty" bson:"date_created,omitempty"`
	Notes        []Note             `json:"notes,omitempty" bson:"notes,omitempty"`
}

type Note struct {
	DateCreated *time.Time `json:"date_created,omitempty" bson:"date_created,omitempty"`
	Msg         string     `json:"msg,omitempty" bson:"msg,omitempty"`
	User        string     `json:"user,omitempty" bson:"user,omitempty"`
}
