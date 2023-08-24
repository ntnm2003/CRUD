package main

import (
	"FiberNewBie/ent"
	user2 "FiberNewBie/ent/user"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
)

func test() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")
	// If you want other engine, just replace with following
	// Create a new engine with django
	// engine := django.New("./views", ".django")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	m := (func(c *fiber.Ctx) error {
		fmt.Println("I'm a middleware")
		return c.Next()
	})
	pages := app.Group("/pages", m)
	pages.Get("/", m, func(c *fiber.Ctx) error {
		// Render index template
		return c.JSON(struct {
			Name string
			Age  int
		}{
			Name: "Robby",
			Age:  20,
		})
	})
	pages.Get("/name/:name/age/:age", func(c *fiber.Ctx) error {
		name := c.Params("name")
		age := c.Params("age")
		return c.Render("index", fiber.Map{
			"Name": name,
			"Age":  age,
		})
	})

	pages.Post("/", m, func(c *fiber.Ctx) error {
		var body struct {
			Message string
		}
		if err := c.BodyParser(&body); err != nil {
			return err
		}

		return c.Render("index", fiber.Map{
			"Name":    "Hello, World!",
			"Message": body.Message,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
func CreateUser(ctx context.Context, client *ent.Client, name string, age int) (*ent.User, error) {
	u, err := client.User.Create().
		SetName(name).SetAge(age).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}
func QueryUserByID(ctx context.Context, client *ent.Client, userID int) (*ent.User, error) {
	user, err := client.User.
		Query().
		Where(user2.IDEQ(userID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func CreateAccount(ctx context.Context, client *ent.Client, a8m *ent.User, username string, password string) (*ent.Account, error) {

	acc, err := client.Account.Create().
		SetUsername(username).SetPassword(password).SetOwner(a8m).Save(ctx)
	if err != nil {
		return nil, err
	}

	return acc, nil
}
