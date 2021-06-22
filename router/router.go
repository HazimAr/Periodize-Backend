package router

import (
	"github.com/NikSchaefer/go-fiber/handlers"
	"github.com/NikSchaefer/go-fiber/middleware"
	"github.com/gofiber/fiber/v2"
)

func Initalize(router *fiber.App) {

	router.Use(middleware.Security)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})

	router.Use(middleware.Json)

	users := router.Group("/users")
	users.Put("/", handlers.CreateUser)
	users.Delete("/", middleware.Authenticated, handlers.DeleteUser)
	users.Patch("/", middleware.Authenticated, handlers.ChangePassword)
	users.Post("/email", handlers.GetUserByEmail)
	users.Patch("/forgot", handlers.ForgotPassword)
	users.Post("/me", middleware.Authenticated, handlers.GetUserInfo)
	users.Post("/login", handlers.Login)
	users.Post("/logout", handlers.Logout)
	users.Patch("/update", middleware.Authenticated, handlers.UpdateUser)

	programs := router.Group("/programs", middleware.Authenticated)
	programs.Put("/", handlers.CreateProgram)
	programs.Post("/all", handlers.GetPrograms)
	programs.Delete("/:id", handlers.DeleteProgram)
	programs.Post("/:id", handlers.GetProgramById)
	programs.Patch("/:id", handlers.UpdateProgram)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}
