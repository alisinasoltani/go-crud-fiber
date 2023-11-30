package routes

import (
	"errors"

	models "github.com/alisinasoltani/goFiber/Models"
	"github.com/alisinasoltani/goFiber/database"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateRespondUser(user models.User) User {
	return User{
		ID: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&user)
	response := CreateRespondUser(user)
	return c.Status(200).JSON(response)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Database.Db.Find(&users)
	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreateRespondUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("User Does Not Exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(200).JSON("Please Ensure That ID is an Integer.")
	}
	var user models.User
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseUser := CreateRespondUser(user)
	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(200).JSON("Please Ensure That ID is an Integer.")
	}
	var user models.User
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
	}
	var updateData UpdateUser
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName
	database.Database.Db.Save(&user)
	resposeUser := CreateRespondUser(user)
	return c.Status(200).JSON(resposeUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(200).JSON("Please Ensure That ID is an Integer.")
	}
	var user models.User
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).SendString("User Successfully Deleted")
}