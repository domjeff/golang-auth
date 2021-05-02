package controller

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	// "strconv"

	"github.com/dgrijalva/jwt-go"
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

	password, _ := bcrypt.GenerateFromPassword([]byte(*data.Password), 1)
	user := &models.User{
		Name:     *data.Name,
		Email:    *data.Email,
		Password: password,
	}
	database.DB.Create(user)
	// c.JSON(user)
	c.JSON(user)
}

func Login(c *fiber.Ctx) {
	var data data

	if err := c.BodyParser(&data); err != nil {
		c.Status(404).Send(err)
		return
	}

	// parameterList := []string{"Name", "Password"}

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(&data)
	json.Unmarshal(inrec, &inInterface)

	// for field, val := range inInterface {
	// 	for _, desiredKey := range parameterList {
	// 		if val == nil && field == desiredKey {
	// 			err := field + " must be filled"
	// 			c.Status(404).Send(err)
	// 			return
	// 		}
	// 	}
	// }

	var user models.User

	database.DB.Where("name= ?", data.Name).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		c.JSON(fiber.Map{
			// "message": "incorrect username and/or password",
			"message": "cannot find",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(*data.Password)); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "incorrect username and/or password",
			},
		)
	}

	c.JSON(user)

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    strconv.Itoa(int(user.Id)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv("secretkey")
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		c.JSON(
			fiber.Map{
				"message": err.Error(),
			},
		)
		return
	}
	cookies := fiber.Cookie{
		Name:     "jwt",
		Value:    ss,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookies)
	c.JSON(
		fiber.Map{
			"message": "success",
		},
	)
	return
}

func Test(c *fiber.Ctx) {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)

	c.JSON(fiber.Map{
		"message": "this is a message",
	})
}
