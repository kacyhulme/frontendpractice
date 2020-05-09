package grifts

import (
	"frontendpractice/models"

	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		u := &models.User{
			FirstName: "Axel",
			LastName:  "Raine",
			Email:     "axelraine@gmail.com",
			City:      "Blue Hill",
			State:     "ME",
		}
		err := models.DB.Create(u)
		return err
	})
})
