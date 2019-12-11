package grifts

import (
	"fmt"
	"github.com/kagundajm/gobuffalo_examples/models"
	"github.com/markbates/grift/grift"
	"github.com/pkg/errors"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add DB seeding stuff here
		if err := models.DB.TruncateAll(); err != nil {
			return errors.WithStack(err)
		}

		categories := models.Categories{
			{Name: "Building Materials"},
			{Name: "Equipment Hire"},
			{Name: "Bakery Products"},
			{Name: "Fasteners"},
		}
		if err := models.DB.Create(&categories); err != nil {
			return errors.WithStack(err)
		}
		fmt.Printf("%+v\n", categories)

		products := models.Products{
			{Name: "Bricks", Price: 25, CategoryID: categories[0].ID},
			{Name: "Doors", Price: 150, CategoryID: categories[0].ID},

			{Name: "Excavator", Price: 2500, CategoryID: categories[1].ID},
			{Name: "Cranes", Price: 1500, CategoryID: categories[1].ID},
			{Name: "Tractors", Price: 2000, CategoryID: categories[1].ID},

			{Name: "Breads", Price: 55, CategoryID: categories[2].ID},
			{Name: "Cakes", Price: 20, CategoryID: categories[2].ID},

			{Name: "Tech Screws", Price: 4, CategoryID: categories[3].ID},
			{Name: "Hex Cap Screws", Price: 7, CategoryID: categories[3].ID},
		}
		if err := models.DB.Create(&products); err != nil {
			return errors.WithStack(err)
		}
		fmt.Printf("%+v\n", products)

		return nil
	})

})
