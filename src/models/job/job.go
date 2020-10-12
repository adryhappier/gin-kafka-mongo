package model

import (
	"time"
)

type Job struct {
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	Company     string    `bson:"company"`
	Salary      string    `bson:"salary"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

type Jobs []Job
