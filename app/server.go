package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
	"github.com/patricksferraz/whoare/domain/service"
	"github.com/patricksferraz/whoare/infra/db"
	"github.com/patricksferraz/whoare/infra/repo"
)

var sessionStore = session.New()

func init() {
	sessionStore.RegisterType(fiber.Map{})
}

func StartApp(orm *db.DbOrm) {
	engine := html.New("./app/views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})
	app.Use(recover.New())
	app.Static("/static", "./public")

	repository := repo.NewRepository(orm)
	service := service.NewService(repository)
	appService := NewAppService(service, sessionStore)

	app.Get("/", appService.Index)
	app.Get("/employees", appService.CreateEmployee)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	err := app.Listen("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start app", err)
	}
}
