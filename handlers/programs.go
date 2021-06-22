package handlers

import (
	"github.com/NikSchaefer/go-fiber/database"
	"github.com/NikSchaefer/go-fiber/model"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Program model.Program

func CreateProgram(c *fiber.Ctx) error {
	db := database.DB
	json := new(Program)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	user := c.Locals("user").(User)
	newProgram := Program{
		ID:			uuid.New(),
		UserRefer:	user.ID,
		Title:      json.Title,
		Descripion:	json.Descripion,
		Days: 		json.Days,
	}

	err := db.Create(&newProgram).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "Creation Error",
		})
	}
	
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}

func GetPrograms(c *fiber.Ctx) error {
	user := c.Locals("user").(User)
	db := database.DB
	Programs := []Program{}
	db.Model(&model.Program{}).Where("user_refer = ?",  user.ID.String()).Order("created_at desc").Limit(100).Find(&Programs)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
		"data":    Programs,
	})
}

func GetProgramById(c *fiber.Ctx) error {
	db := database.DB
	param := c.Params("id")
	id, err := uuid.Parse(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID Format",
		})
	}
	program := Program{}
	query := Program{ID: id}
	err = db.First(&program, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Program not found",
		})
	}
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
		"data":    program,
	})
}

func UpdateProgram(c *fiber.Ctx) error {
	type UpdateProgramRequest struct {
		Name      	string 		`json:"name"`
		Descripion 	string 		`json:"description"`
		Private		bool 		`json:"private"`
		Experience  string		`json:"experience"`
		Sport		string		`json:"sport"`
		Sessionid 	string 		`json:"sessionid"`
	}

	db := database.DB
	user := c.Locals("user").(User)
	json := new(UpdateProgramRequest)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	param := c.Params("id")
	id, err := uuid.Parse(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID format",
		})
	}

	found := Program{}
	query := Program{
		ID:        id,
		UserRefer: user.ID,
	}

	err = db.First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Program not found",
		})
	}

	// found.Name = json.Name
	// found.Descripion = json.Descripion
	// found.Private = json.Private
	// found.Experience = json.Experience
	// found.Sport = json.Sport
	
	db.Save(&found)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}

func DeleteProgram(c *fiber.Ctx) error {
	db := database.DB
	user := c.Locals("user").(User)
	param := c.Params("id")
	id, err := uuid.Parse(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID format",
		})
	}
	found := Program{}
	query := Program{
		ID:        id,
		UserRefer: user.ID,
	}
	err = db.First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Program not found",
		})
	}
	db.Delete(&found)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}
