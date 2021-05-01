package controller

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber"
)

func Hello(c *fiber.Ctx) {
	c.SendString("Hello Worldaaa")
}

type data struct {
	Name     *string `json:"name" xml:"name" form:"name"`
	Password *string `json:"password" xml:"password" form:"password"`
	Token    *string `json:"token" xml:"token" form:"token"`
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
		fmt.Println(val)
		if val == nil {
			err := field + " must be filled"
			c.Status(404).Send(err)
			return
		}
	}

	c.JSON(data)

}
