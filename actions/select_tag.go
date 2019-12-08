package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/kagundajm/gobuffalo_examples/models"
	"github.com/pkg/errors"
)

const txError = "Unable to get a database transaction"
const subTitle = "Select Tags Example"

// ListProducts gets all products
func ListProducts(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New(txError))
	}

	products := []models.ProductListModel{}
	s := "select p.id, p.name, p.price, p.category_id, c.name as category_name from products p join categories c on c.id = p.category_id ORDER BY c.name, p.name;"

	if err := tx.RawQuery(s).All(&products); err != nil {
		return errors.WithStack(err)
	}

	c.Set("sub-title", subTitle)
	c.Set("products", products)
	return c.Render(200, r.HTML("select_tags/index.html"))
}

// NewProduct renders form for creating product
func NewProduct(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New(txError))
	}

	categories := models.Categories{}
	if err := tx.Order("name").All(&categories); err != nil {
		return errors.WithStack(err)
	}

	product := &models.Product{}

	c.Set("sub-title", subTitle)
	c.Set("product", product)
	c.Set("categories", categories)
	return c.Render(200, r.HTML("select_tags/new.html"))
}

// EditProduct renders product edit form
func EditProduct(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New(txError))
	}

	product := &models.Product{}
	if err := tx.Find(product, c.Param("product_id")); err != nil {
		return c.Error(404, err)
	}

	categories := models.Categories{}
	if err := tx.Order("name").All(&categories); err != nil {
		return errors.WithStack(err)
	}

	c.Set("sub-title", subTitle)
	c.Set("product", product)
	c.Set("categories", categories)
	return c.Render(200, r.HTML("select_tags/edit.html"))
}

func SaveProduct(c buffalo.Context) error {
	// Not implemented
	return c.Redirect(302, "selectTagsPath()")
}

func UpdateProduct(c buffalo.Context) error {
	// Not implemented
	return c.Redirect(302, "selectTagsPath()")
}
