package routes

import (
	"os"
	"time"

	"url-shortener/database"
	"url-shortener/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"shortcode"`
	Expiry      time.Duration `json:"validity"`
}

type response struct {
	CustomShort string        `json:"shortLink"`
	Expiry      time.Duration `json:"expiry"`
}

func ShortenURL(c *fiber.Ctx) error {

	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Cheeky huh?"})
	}

	body.URL = helpers.EnforceHTTP(body.URL)

	var id string

	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	r := database.CreateClient(0)
	defer r.Close()

	val, _ := r.Get(id).Result()
	if val != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Short already exists!"})
	}

	if body.Expiry == 0 {
		body.Expiry = 30
	}

	err := r.Set(id, body.URL, time.Duration(body.Expiry*time.Minute)).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot connect to Database!"})
	}

	resp := response{
		CustomShort: os.Getenv("DOMAIN") + "/" + id,
		Expiry:      body.Expiry,
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
