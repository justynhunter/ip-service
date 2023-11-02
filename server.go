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
  re := regexp.MustCompile(`(^127\.)|(^10\.)|(^172\.1[6-9]\.)|(^172\.2[0-9]\.)|(^172\.3[0-1]\.)|(^192\.168\.)`)
  if re.MatchString(ip) {
  	ips := c.IPs()
  	if (len(ips) > 0) {
      for _, i := range ips {
        ip = ip + "\n" + i
      }
  	}
  }

  return c.SendString(ip + "\n")
}