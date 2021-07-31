package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"fmt"
)

type Reservation struct {
	Reservee string `json:"reservee`
	Date	 string  `json:"date"`
	Time     string  `json:"time"`
	Court 	 string  `json:"court"`
}

type Court struct {
	Name 	 string `json:"name"`
	Reserved []Reservation `json:"reserved"`
	Location string `json:"location"`
}

var reservations = []Reservation{
	{Reservee: "None", Date: "None", Time: "None", Court: "None"},
}

var courts = []Court{
	{Name: "One", Reserved: reservations, Location: "Deez nuts ave"},
	{Name: "Two", Reserved: reservations, Location: "Deez nuts lane"},
	{Name: "Three", Reserved: reservations, Location: "Deez nuts road"},
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
		reservation := new(Reservation)

		if err := ctx.BodyParser(reservation); err != nil {
			fmt.Println("error = ", err)
			return ctx.SendStatus(200)
		}

		updateReservations(reservation);
		return nil
	})

	app.Listen(":3000")
}

func updateReservations(class *Reservation){
	for  i:=0; i < len(courts); i++ {
		if ( courts[i].Name == class.Court ){
			fmt.Println("Found court!");
			courts[i].Reserved = append(courts[i].Reserved, *class);
			fmt.Println(courts);
		}
	}
}
