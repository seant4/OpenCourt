package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Reservation struct {
	Reservee string `json:"reservee`
	Date	 string  `json:"date"`
	Time     string  `json:"time"`
}

type Court struct {
	Name 	 string `json:"name"`
	Reserved Reservation `json:"reserved"`
	Location string `json:"location"`
}

var courts = []Court{
	{Name: "One", Reserved: Reservation{Reservee: "Jon", Date: "07-28-2021", Time: "7:00AM"}, Location: "Deez nuts ave"},
	{Name: "Two", Reserved: Reservation{Reservee: "Tod", Date: "07-29-2021", Time: "8:00AM"}, Location: "Deez nuts lane"},
	{Name: "Three", Reserved: Reservation{Reservee: "Tim", Date: "07-29-2021", Time: "9:00AM"}, Location: "Deez nuts road"},

}

func main(){

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		Next:             nil,
    	AllowOrigins:     "*",
    	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
    	AllowHeaders:     "",
    	AllowCredentials: false,
    	ExposeHeaders:    "",
    	MaxAge:           0,
	}))

	app.Static("/", "./public")


	app.Get("/courts", func(ctx *fiber.Ctx) error{
		ctx.Status(fiber.StatusOK).JSON(courts)
		return nil
	})

	app.Post("/courts", func(ctx *fiber.Ctx) error {
		payload := struct {
			Reservee string `json:"reservee"`
			Date 	 string `json:"Date"`
			Time     string `json:"Time"`
		}{}

		if err := ctx.BodyParser(&payload); err != nil {
			return err  
		}
		return ctx.JSON(payload);
	})

	app.Listen(":3000")
}

