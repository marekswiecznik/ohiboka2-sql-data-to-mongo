package mongo

import (
	"time"
)

type Bracelet struct {
	ID         int       `json:"id"`
	Created    time.Time `json:"created"`
	Name       string    `json:"name"`
	Accepted   bool      `json:"accepted"`
	Difficulty int       `json:"difficulty"`
	Category   *string   `json:"category"`
	Rate       string    `json:"rate"`
	Public     bool      `json:"public"`
	Slug       string    `json:"slug"`
	Deleted    bool      `json:"deleted"`
	Type       string    `json:"type"`

	Photo   *Photo   `json:"photo,omitempty"`
	Author  User     `json:"author"`
	Photos  []Photo  `json:"photos"`
	Rates   []Rate   `json:"rates"`
	Strings []string `json:"strings"`
	Rows    []Row    `json:"rows"`
}

type Photo struct {
	Filename string `json:"filename"`
	Author   User   `json:"author"`
	Accepted bool   `json:"accepted"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Rate struct {
	Rate int  `json:"rate"`
	User User `json:"user"`
}

type Row struct {
	Odd   bool   `json:"odd"`
	Knots []Knot `json:"knots"`
}

type Knot struct {
	Type string `json:"type"`
}
