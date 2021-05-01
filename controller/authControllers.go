package controller

import (
	"encoding/json"

	"github.com/domjeff/golang-auth/database"
	"github.com/domjeff/golang-auth/models"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

func Hello(c *fiber.Ctx) {
	c.SendString("Hello Worldaaa")
}

type data struct {
	Name     *string `json:"name" xml:"name" form:"name"`
	Password *string `json:"password" xml:"password" form:"password"`
	Email    *string `json:"email" xml:"email" form:"email"`
	// Token    *string `json:"token" xml:"token" form:"token"`
}

func Register(c *fiber.Ctx) {
	var data data

	if err := c.BodyParser(&data); err != nil {
		c.Status(404).Send(err)
		return
	}

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(&data)
	json.Unmarshal(inrec, &inInterface)
	for field, val := range inInterface {
		// fmt.Println(val)
		if val == nil {
			err := field + " must be filled"
			c.Status(404).Send(err)
			return
		}
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(*data.Password), 14)
	user := &models.User{
		Name:     *data.Name,
		Email:    *data.Email,
		Password: password,
	}
	database.DB.Create(user)
	// c.JSON(user)
	c.JSON(user)
}
