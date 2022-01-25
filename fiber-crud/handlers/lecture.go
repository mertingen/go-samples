package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mertingen/go-samples/config"
	"github.com/mertingen/go-samples/entities"
)

type lectureResp struct {
	Status  bool               `json:"status"`
	Message string             `json:"message"`
	Data    []entities.Lecture `json:"data"`
}

func GetLectures(c *fiber.Ctx) error {
	rows := make([]entities.Lecture, 0)

	result := config.Database.Find(&rows)

	resp := lectureResp{
		Status:  true,
		Message: "Lectures are found.",
		Data:    rows,
	}
	if result.RowsAffected == 0 {
		resp.Status = false
		resp.Message = "Lectures are not found."
		return c.Status(200).JSON(resp)
	}
	return c.Status(200).JSON(resp)
}

func GetLecture(c *fiber.Ctx) error {
	id := c.Params("id")
	row := make([]entities.Lecture, 0)

	result := config.Database.Find(&row, id)

	resp := lectureResp{
		Status:  true,
		Message: "Lecture is found.",
		Data:    row,
	}
	if result.RowsAffected == 0 {
		resp.Status = false
		resp.Message = "Lecture is not found."
		return c.Status(200).JSON(resp)
	}

	return c.Status(200).JSON(&resp)
}

func AddLecture(c *fiber.Ctx) error {
	row := new(entities.Lecture)

	resp := lectureResp{
		Status:  true,
		Message: "Lecture is created.",
		Data:    make([]entities.Lecture, 0),
	}
	if err := c.BodyParser(row); err != nil {
		resp.Status = false
		resp.Message = "An error occurs while lecture is created."
		return c.Status(200).JSON(resp)
	}

	result := config.Database.Create(&row)
	if result.RowsAffected == 0 {
		resp.Status = false
		resp.Message = "An error occurs while lecture is created."
		return c.Status(200).JSON(resp)
	}

	resp.Data = append(resp.Data, *row)
	return c.Status(201).JSON(resp)
}

func UpdateLecture(c *fiber.Ctx) error {
	row := new(entities.Lecture)
	id := c.Params("id")

	resp := lectureResp{
		Status:  true,
		Message: "Lecture is updated.",
		Data:    make([]entities.Lecture, 0),
	}
	if err := c.BodyParser(row); err != nil {
		resp.Status = false
		resp.Message = "An error occurs while lecture is updated."
		return c.Status(200).JSON(resp)
	}

	result := config.Database.Where("id = ?", id).Updates(&row)
	if result.RowsAffected == 0 {
		resp.Status = false
		resp.Message = "An error occurs while lecture is updated."
		return c.Status(200).JSON(resp)
	}

	resp.Data = append(resp.Data, *row)
	return c.Status(200).JSON(resp)
}

func RemoveLecture(c *fiber.Ctx) error {
	id := c.Params("id")
	var row entities.Lecture

	resp := lectureResp{
		Status:  true,
		Message: "Lecture is deleted.",
		Data:    make([]entities.Lecture, 0),
	}
	result := config.Database.Delete(&row, id)

	if result.RowsAffected == 0 {
		resp.Status = false
		resp.Message = "An error occurs while lecture is updated."
		return c.Status(200).JSON(resp)
	}

	return c.Status(200).JSON(resp)
}
