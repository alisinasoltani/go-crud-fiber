package routes

import (
	"errors"

	models "github.com/alisinasoltani/goFiber/Models"
	"github.com/alisinasoltani/goFiber/database"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(prodctModel models.Product) Product {
	return Product{
		ID: prodctModel.ID,
		Name:  prodctModel.Name,
		SerialNumber: prodctModel.SerialNumber,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.Database.Db.Find(&products)
	responseProducts := []Product{}
	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}
	return c.Status(200).JSON(responseProducts)
}

func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("User Does Not Exist")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(200).JSON("Please Ensure That ID is an Integer.")
	}
	var product models.Product
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(200).JSON("Please Ensure That ID is an Integer.")
	}
	var product models.Product
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateProduct struct {
		Name string `json:"first_name"`
		SerialNumber string `json:"last_name"`
	}
	var updateData UpdateProduct
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	product.Name = updateData.Name
	product.SerialNumber = updateData.SerialNumber
	database.Database.Db.Save(&product)
	resposeProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(resposeProduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(200).JSON("Please Ensure That ID is an Integer.")
	}
	var product models.Product
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).SendString("Product Successfully Deleted")
}