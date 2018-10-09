package models

import (
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Product struct {
	ID         uuid.UUID `json:"id" db:"id"`
	CategoryID uuid.UUID `json:"category_id" db:"category_id"`
	Name       string    `json:"name" db:"name"`
	Price      int       `json:"price" db:"price"`
}

type Products []Product
type ProductListModel struct {
	Product
	CategoryName string `json:"category_name" db:"category_name"`
}

// Validate product has a name, category ID and price
func (p *Product) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Name, Name: "Name"},
		&validators.UUIDIsPresent{Field: p.CategoryID, Name: "CategoryID"},
		&validators.IntIsGreaterThan{Field: p.Price, Name: "Price", Compared: 1},
	), nil
}
