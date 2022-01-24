package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mertingen/go-samples/config"
	"github.com/mertingen/go-samples/entities"
)

func GetStudents(c *fiber.Ctx) error {
	var student []entities.Student

	config.Database.Find(&student)
	return c.Status(200).JSON(student)
}

func GetStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	var student entities.Student

	result := config.Database.Find(&student, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&student)
}

func AddStudent(c *fiber.Ctx) error {
	student := new(entities.Student)

	if err := c.BodyParser(student); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&student)
	return c.Status(201).JSON(student)
}

func UpdateStudent(c *fiber.Ctx) error {
	student := new(entities.Student)
	id := c.Params("id")

	if err := c.BodyParser(student); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&student)
	return c.Status(200).JSON(student)
}

func RemoveStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	var student entities.Student

	result := config.Database.Delete(&student, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
