package server

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gituhub.com/snirkop89/goly/app/model"
	"gituhub.com/snirkop89/goly/utils"
)

func redirect(c *fiber.Ctx) error {
	golyURL := c.Params("redirect")
	goly, err := model.FindByGolyURL(golyURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not find goly in DB " + err.Error(),
		})
	}

	// grab any stats you want
	goly.Clicked += 1
	err = model.UpdateGoly(goly)
	if err != nil {
		fmt.Printf("error updating stat")
	}

	return c.Redirect(goly.Redirect, fiber.StatusTemporaryRedirect)
}

func getAllGolies(c *fiber.Ctx) error {
	golies, err := model.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all goly links " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(golies)
}

func getGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id " + err.Error(),
		})
	}

	goly, err := model.GetOne(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retrieve goly " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}

func createGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var goly model.Goly
	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}

	if goly.Random {
		goly.Goly = utils.RandomURL(8)
	}

	err = model.CreateGoly(goly)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error creating goly " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}

func updateGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var goly model.Goly
	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}

	err = model.UpdateGoly(goly)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error updating goly " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}

func deleteGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error could not parse id " + err.Error(),
		})
	}

	err = model.DeleteGoly(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error deleting goly " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "goly deleted",
	})
}

func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/:redirect", redirect)

	router.Get("/goly", getAllGolies)
	router.Get("/goly/:id", getGoly)
	router.Post("/goly", createGoly)
	router.Patch("goly", updateGoly)
	router.Delete("/goly/:id", deleteGoly)

	router.Listen(":8888")
}
