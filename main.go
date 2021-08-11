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
	{Name: "One", Reserved: reservations, Location: "Mecklenburg"},
	{Name: "Two", Reserved: reservations, Location: "Joe mama"},
	{Name: "Three", Reserved: reservations, Location: "Mecklenburg"},
	{Name: "Four", Reserved: reservations, Location: "Joe mama"},
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

		result := updateReservations(reservation);
		if(!result){
			return ctx.Status(200).SendString("Reservation successfully made!");
			
		}else{
			return ctx.Status(406).SendString("Reservation is taken, please try a different date or time!");

		}
		return nil
	})

	app.Listen(":3000")
}

func updateReservations(class *Reservation) (bool){
	duplicate := false;
	for  i:=0; i < len(courts); i++ {
		if ( courts[i].Name == class.Court ){
			fmt.Println("Found court!");
			for j := 0; j < len(courts[i].Reserved); j++ {
				if(courts[i].Reserved[j].Date == class.Date && courts[i].Reserved[j].Time == class.Time){
					fmt.Println("Duplicate");
					duplicate = true;
				}
			}
			if(!duplicate){
				courts[i].Reserved = append(courts[i].Reserved, *class);
				fmt.Println(courts);
				return duplicate;
			}else{
				return duplicate;
			}
		}
	}
	return duplicate
}
