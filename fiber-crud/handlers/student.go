package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mertingen/go-samples/config"
	"github.com/mertingen/go-samples/entities"
)

type response struct {
	Status  bool               `json:"status"`
	Message string             `json:"message"`
	Data    []entities.Student `json:"data"`
}

func GetStudents(c *fiber.Ctx) error {
	students := make([]entities.Student, 0)

	result := config.Database.Find(&students)

	resp := response{
		Status:  true,
		Message: "Students are found.",
		Data:    students,
	}
	if result.RowsAffected == 0 {
		resp.Status = false
		resp.Message = "Students are not found."
		return c.Status(200).JSON(resp)
	}
	return c.Status(200).JSON(resp)
}

func GetStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	student := make([]entities.Student, 0)

	result := config.Database.Find(&student, id)

	resp := response{
		Status:  true,
		Message: "Student is found.",
		Data:    student,
	}
	if result.RowsAffected == 0 {
		resp.Status = false
		resp.Message = "Students is not found."
		return c.Status(200).JSON(resp)
	}

	return c.Status(200).JSON(&resp)
}

func AddStudent(c *fiber.Ctx) error {
	student := new(entities.Student)

	resp := response{
		Status:  true,
		Message: "Student is created.",
		Data:    make([]entities.Student, 0),
	}
	if err := c.BodyParser(student); err != nil {
		resp.Status = false
		resp.Message = "An error occurs while player is created."
		return c.Status(200).JSON(resp)
	}

	result := config.Database.Create(&student)
	if result.RowsAffected == 0 {
		resp.Status = false
		resp.Message = "An error occurs while player is created."
		return c.Status(200).JSON(resp)
	}

	resp.Data = append(resp.Data, *student)
	return c.Status(201).JSON(resp)
}

func UpdateStudent(c *fiber.Ctx) error {
	student := new(entities.Student)
	id := c.Params("id")

	resp := response{
		Status:  true,
		Message: "Student is updated.",
		Data:    make([]entities.Student, 0),
	}
	if err := c.BodyParser(student); err != nil {
		resp.Status = false
		resp.Message = "An error occurs while player is updated."
		return c.Status(200).JSON(resp)
	}

	result := config.Database.Where("id = ?", id).Updates(&student)
	if result.RowsAffected == 0 {
		resp.Status = false
		resp.Message = "An error occurs while player is updated."
		return c.Status(200).JSON(resp)
	}

	resp.Data = append(resp.Data, *student)
	return c.Status(200).JSON(resp)
}

func RemoveStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	var student entities.Student

	resp := response{
		Status:  true,
		Message: "Student is deleted.",
		Data:    make([]entities.Student, 0),
	}
	result := config.Database.Delete(&student, id)

	if result.RowsAffected == 0 {
		resp.Status = false
		resp.Message = "An error occurs while player is updated."
		return c.Status(200).JSON(resp)
	}

	return c.Status(200).JSON(resp)
}
