package models

import (
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Category struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
}

type Categories []Category

// Validate we have a name for the category
func (c *Category) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return nil, validate.Validate(
		&validators.StringIsPresent{Field: c.Name, Name: "Name"},
	)
}

// SelectLabel and SelectValue implement Selectable interface which will
// allow passing of Categories as select options
// https://github.com/gobuffalo/tags/wiki/Select-Tag#selectable-interface

// SelectLabel - label for select tag options
func (c Category) SelectLabel() string {
	return c.Name
}

// SelectValue - value for select tag options
func (c Category) SelectValue() interface{} {
	return c.ID
}
