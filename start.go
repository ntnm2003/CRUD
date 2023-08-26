package main

import (
	"FiberNewBie/ent"
	"FiberNewBie/ent/account"
	user2 "FiberNewBie/ent/user"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

func main() {

	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=iam password=222003 sslmode=disable")

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	ctx := context.Background()
	engine := html.New("./views", ".html")
	// If you want other engine, just replace with following
	// Create a new engine with django
	// engine := django.New("./views", ".django")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Post("/user/create/:name/:age", func(c *fiber.Ctx) error {
		name := c.Params("name")
		age, _ := strconv.Atoi(c.Params("age"))

		user, err := CreateUser(ctx, client, name, age)
		if err != nil {
			log.Fatalf("error querying users: %v", err)
		}
		return c.JSON(user)
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		user, err := client.Account.Query().Where(account.Username(username)).Only(ctx)
		if err != nil {
			return c.SendStatus(http.StatusUnauthorized)
		}

		// Verify the password (you should use a proper password hashing library)
		if user.Password != password {
			return c.SendStatus(http.StatusUnauthorized)
		}

		return c.SendString(fmt.Sprintf("Welcome, %s!", user.Username))
	})

	app.Get("/user/list", func(c *fiber.Ctx) error {
		users, err := client.User.Query().All(ctx)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error querying users")
		}
		return c.JSON(users)
	})

	app.Get("/user/:uid", func(c *fiber.Ctx) error {

		uid, _ := strconv.Atoi(c.Params("uid"))
		user, err := QueryUserByID(ctx, client, uid)
		if err != nil {
			log.Fatalf("wrong")
		}
		return c.JSON(user)
	})
	app.Get("/acc/list", func(c *fiber.Ctx) error {
		users, err := client.Account.Query().Where(account.IDEQ(2)).All(ctx)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error querying users")
		}
		return c.JSON(users)
	})
	app.Get("/acc/user/:uid", func(c *fiber.Ctx) error {
		uid, _ := strconv.Atoi(c.Params("uid"))
		acc, err := client.User.Query().Where(user2.IDEQ(uid)).QueryAccount().Only(ctx)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error querying users")
		}

		return c.JSON(acc)

	})
	app.Get("/acc/create/:uid/:username/:password", func(c *fiber.Ctx) error {

		uid, _ := strconv.Atoi(c.Params("uid"))
		username := c.Params("username")
		password := c.Params("password")
		user, err1 := QueryUserByID(ctx, client, uid)
		if err1 != nil {
			log.Fatalf("can not find the user")
		}
		acc, err := CreateAccount(ctx, client, user, username, password)
		if err != nil {
			log.Fatalf("fault")
		}
		return c.JSON(acc)
	})
	log.Fatal(app.Listen(":3000"))
}
