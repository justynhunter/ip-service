package main

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// routes
	app.Get("/", getIP)

	app.Listen(":3000")
}

func getIP(c *fiber.Ctx) error {
	ip := c.IP()

	/* if we get a reserved ip we want to see if there
	 * are ips in the headers, if so use the first one */
	re := regexp.MustCompile("")
	if re.MatchString(ip) {
		ips := c.IPs()
		if (len(ips) > 0) {
			ip = ips[0]
		}
	}

	return c.SendString(ip)
}