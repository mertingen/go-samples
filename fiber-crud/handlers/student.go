package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mertingen/go-samples/config"
	"github.com/mertingen/go-samples/entities"
)

type studentResp struct {
	Status  bool               `json:"status"`
	Message string             `json:"message"`
	Data    []entities.Student `json:"data"`
}

func GetStudents(c *fiber.Ctx) error {
	rows := make([]entities.Student, 0)

	result := config.Database.Find(&rows)

	resp := studentResp{
		Status:  true,
		Message: "Students are found.",
		Data:    rows,
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
	row := make([]entities.Student, 0)

	result := config.Database.Find(&row, id)

	resp := studentResp{
		Status:  true,
		Message: "Student is found.",
		Data:    row,
	}
	if result.RowsAffected == 0 {
		resp.Status = false
		resp.Message = "Students is not found."
		return c.Status(200).JSON(resp)
	}

	return c.Status(200).JSON(&resp)
}

func AddStudent(c *fiber.Ctx) error {
	row := new(entities.Student)

	resp := studentResp{
		Status:  true,
		Message: "Student is created.",
		Data:    make([]entities.Student, 0),
	}
	if err := c.BodyParser(row); err != nil {
		resp.Status = false
		resp.Message = "An error occurs while student is created."
		return c.Status(200).JSON(resp)
	}

	result := config.Database.Create(&row)
	if result.RowsAffected == 0 {
		resp.Status = false
		resp.Message = "An error occurs while student is created."
		return c.Status(200).JSON(resp)
	}

	resp.Data = append(resp.Data, *row)
	return c.Status(201).JSON(resp)
}

func UpdateStudent(c *fiber.Ctx) error {
	row := new(entities.Student)
	id := c.Params("id")

	resp := studentResp{
		Status:  true,
		Message: "Student is updated.",
		Data:    make([]entities.Student, 0),
	}
	if err := c.BodyParser(row); err != nil {
		resp.Status = false
		resp.Message = "An error occurs while student is updated."
		return c.Status(200).JSON(resp)
	}

	result := config.Database.Where("id = ?", id).Updates(&row)
	if result.RowsAffected == 0 {
		resp.Status = false
		resp.Message = "An error occurs while student is updated."
		return c.Status(200).JSON(resp)
	}

	resp.Data = append(resp.Data, *row)
	return c.Status(200).JSON(resp)
}

func RemoveStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	var row entities.Student

	resp := studentResp{
		Status:  true,
		Message: "Student is deleted.",
		Data:    make([]entities.Student, 0),
	}
	result := config.Database.Delete(&row, id)

	if result.RowsAffected == 0 {
		resp.Status = false
		resp.Message = "An error occurs while student is updated."
		return c.Status(200).JSON(resp)
	}

	return c.Status(200).JSON(resp)
}
