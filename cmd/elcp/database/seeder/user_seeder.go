package seeder

import (
	"github.com/riipandi/elisacp/cmd/elcp/model"
	"github.com/riipandi/elisacp/cmd/elcp/utils"
)

var password, _ = utils.HashPassword("secret")

var users = []model.User{
	{
		Name:     "Admin Sistem",
		Email:    "admin@example.com",
		Username: "admin",
		Password: password,
		IsActive: 1,
		},
	model.User{
		Name: "John Doe",
		Email: "johndoe@gmail.com",
		Username: "johndoe",
		Password: password,
		IsActive: 0,
	},
}
