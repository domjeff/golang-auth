package controller

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	// "strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/domjeff/golang-auth/cache"
	"github.com/domjeff/golang-auth/database"
	"github.com/domjeff/golang-auth/models"

	// "github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello Worldaaa")
}

type data struct {
	Name     *string `json:"name" xml:"name" form:"name"`
	Password *string `json:"password" xml:"password" form:"password"`
	Email    *string `json:"email" xml:"email" form:"email"`
	// Token    *string `json:"token" xml:"token" form:"token"`
}

func Register(c *fiber.Ctx) error {
	var data data

	if err := c.BodyParser(&data); err != nil {
		// c.Status(404).Send(err)
		return fiber.NewError(404, err.Error())
	}

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(&data)
	json.Unmarshal(inrec, &inInterface)
	for field, val := range inInterface {
		// fmt.Println(val)
		if val == nil {
			err := field + " must be filled"
			return fiber.NewError(404, err)
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
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data data

	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(404, err.Error())
	}

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(&data)
	json.Unmarshal(inrec, &inInterface)

	var user models.User

	database.DB.Where("name= ?", data.Name).First(&user)

	if user.Id == 0 {
		return c.
			Status(fiber.StatusNotFound).
			JSON(fiber.Map{
				"message": "cannot find",
			})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(*data.Password)); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "incorrect username and/or password",
			},
		)
	}

	ss, err := GenerateJwtToken(user)
	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(
				fiber.Map{
					"message": err.Error(),
				},
			)
	}

	userCache := cache.SetupUserCache()
	if err = userCache.CheckUserToken(user); err != nil {
		return c.Status(fiber.StatusMethodNotAllowed).
			JSON(
				fiber.Map{
					"message": err.Error(),
				},
			)
	}

	if err = userCache.SetUserToken(user, *ss); err != nil {
		return c.Status(fiber.StatusMethodNotAllowed).
			JSON(
				fiber.Map{
					"message": err.Error(),
				},
			)
	}

	cookies := fiber.Cookie{
		Name:     "jwt",
		Value:    *ss,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookies)
	return c.JSON(
		fiber.Map{
			"message": "success",
		},
	)

}

func User(c *fiber.Ctx) error {
	cookies := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookies, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		secretKey := os.Getenv("secretKey")
		return []byte(secretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookies := fiber.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	}

	c.Cookie(&cookies)
	return c.JSON(
		fiber.Map{
			"message": "log out success",
		},
	)
}

func GenerateJwtToken(user models.User) (*string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    strconv.Itoa(int(user.Id)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv("secretkey")
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	var res string
	res = ss
	return &res, nil
}
