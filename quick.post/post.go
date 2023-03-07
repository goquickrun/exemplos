package main

import "github.com/jeffotoni/quick"

type My struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}

func main() {
	app := quick.New()
	app.Post("/v1/user", func(c *quick.Ctx) error {
		var my My
		err := c.BodyParser(&my)
		if err != nil {
			c.Status(400).SendString(err.Error())
		}

		return c.Status(200).JSON(&my)
		// ou
		//c.Status(200).String(c.BodyString())
	})

	app.Listen("0.0.0.0:8080")
}
