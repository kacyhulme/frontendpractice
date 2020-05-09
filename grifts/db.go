package grifts

import (
	"frontendpractice/models"

	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		u := &models.Users{
			{FirstName: "Daddy",
				LastName: "Pants",
				Email:    "daddy@gmail.com",
				City:     "Blue Hill",
				State:    "ME"},
			{FirstName: "Detroit",
				LastName: "Pants",
				Email:    "detroit@gmail.com",
				City:     "Blue Hill",
				State:    "ME"},
		}
		err := models.DB.Create(u)
		return err
	})
})
