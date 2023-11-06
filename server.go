package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// routes
	app.Get("/", GetIP)
	app.Get("/ip", GetIP)
	app.Get("/ns", GetIpForDomain)

	err := app.Listen(":3000")
	println(err.Error())
}

func GetIpForDomain(c *fiber.Ctx) error {
	domain := c.Query("n")
	result := NsLookup(domain)
	return c.JSON(result)
}

func GetIP(c *fiber.Ctx) error {
  result := GetClientIp(c.IP(), c.IPs())
  return c.JSON(result)
}