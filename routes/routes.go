package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kibo/e-wallet/controllers"
)

func Setup(app *fiber.App)  {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/profile", controllers.Profile)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/users", controllers.Users)

	app.Get("/api/history/:id", controllers.History)
	app.Post("/api/topup/:id", controllers.Topup)
	app.Post("/api/transfer/:id", controllers.Transfer)
	
}