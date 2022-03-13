package front

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
	sessionStore.RegisterType(Session{})
}

func StartFront(orm *db.DbOrm) {
	engine := html.New("./app/front/views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})
	app.Use(recover.New())
	app.Static("/static", "./app/front/public")

	repository := repo.NewRepository(orm)
	service := service.NewService(repository)
	front := NewFront(service, sessionStore)
	// middleware := NewMiddleware(sessionStore)

	app.Get("/", front.Index)
	app.Get("/register", front.Register)
	app.Post("/register", front.Register)
	app.Get("/profile/:employee_id", front.Profile)
	app.Get("/profile/:employee_id/edit", front.ProfileEdit)
	app.Post("/profile/:employee_id/edit", front.ProfileEdit)

	app.Use(func(c *fiber.Ctx) error {
		return c.Render("errors/404", fiber.Map{"Status": 404, "Error": "Not Found"}) // => 404 "Not Found"
	})

	err := app.Listen("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start app", err)
	}
}
