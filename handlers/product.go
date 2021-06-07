package handlers

import (
	"strconv"

	"github.com/NikSchaefer/go-fiber/database"
	"github.com/NikSchaefer/go-fiber/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateProduct(c *fiber.Ctx) error {
	db := database.DB
	json := new(Product)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	user := c.Locals("user").(User)
	newProduct := Product{
		UserRefer:	user.ID,
		Name:      	json.Name,
		Descripion:	json.Descripion,
		Experience: json.Experience,
		Private:	json.Private,
		Sport:		json.Sport,
	}

	err := db.Create(&newProduct).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}

func GetProducts(c *fiber.Ctx) error {
	db := database.DB
	Products := []Product{}
	db.Model(&model.Product{}).Order("ID asc").Limit(100).Find(&Products)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
		"data":    Products,
	})
}

func GetProductById(c *fiber.Ctx) error {
	db := database.DB
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID Format",
		})
	}
	product := Product{}
	query := Product{ID: id}
	err = db.First(&product, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Product not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	type UpdateProductRequest struct {
		Name      	string 		`json:"name"`
		Descripion 	string 		`json:"description"`
		Private		bool 		`json:"private"`
		Experience  string		`json:"experience"`
		Sport		string		`json:"sport"`
		Sessionid 	string 		`json:"sessionid"`
	}

	db := database.DB
	user := c.Locals("user").(User)
	json := new(UpdateProductRequest)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID format",
		})
	}

	found := Product{}
	query := Product{
		ID:        id,
		UserRefer: user.ID,
	}

	err = db.First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Product not found",
		})
	}

	found.Name = json.Name
	found.Descripion = json.Descripion
	found.Private = json.Private
	found.Experience = json.Experience
	found.Sport = json.Sport
	
	db.Save(&found)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	db := database.DB
	user := c.Locals("user").(User)
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID format",
		})
	}
	found := Product{}
	query := Product{
		ID:        id,
		UserRefer: user.ID,
	}
	err = db.First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Product not found",
		})
	}
	db.Delete(&found)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}
