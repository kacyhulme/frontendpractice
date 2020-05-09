package grifts

import (
	"frontendpractice/models"

	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		u := &models.Users{
			{FirstName: "Luke",
				LastName: "Wendling",
				Email:    "luke@gmail.com",
				City:     "Blue Hill",
				State:    "ME"},
			{FirstName: "Detroit",
				LastName: "Wendling",
				Email:    "detroit@gmail.com",
				City:     "Blue Hill",
				State:    "ME"},
			{FirstName: "Greta",
				LastName: "Hulme",
				Email:    "thisisgreta@gmail.com",
				City:     "Blue Hill",
				State:    "ME"},
			{FirstName: "Axel",
				LastName: "Raine",
				Email:    "arw@gmail.com",
				City:     "Blue Hill",
				State:    "ME"},
		}
		err := models.DB.Create(u)
		return err
	})
})
