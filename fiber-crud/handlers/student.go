package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mertingen/go-samples/config"
	"github.com/mertingen/go-samples/entities"
	"log"
)

type studentResp struct {
	Status  bool               `json:"status"`
	Message string             `json:"message"`
	Data    []entities.Student `json:"data"`
}

func GetStudents(c *fiber.Ctx) error {
	rows := make([]entities.Student, 0)

	result := config.Database.Preload("Lectures").Find(&rows)

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

	result := config.Database.Preload("Lectures").Find(&row, id)

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

func AttachLectures(c *fiber.Ctx) error {
	id := c.Params("id")
	student := make([]entities.Student, 0)
	lectures := make([]entities.Lecture, 0)

	result := config.Database.Find(&student, id)

	if result.RowsAffected == 0 {
		resp := studentResp{
			Status:  false,
			Message: "Students are not found.",
			Data:    student,
		}
		return c.Status(200).JSON(resp)
	}

	lectureIds := new([]uint)
	if err := c.BodyParser(lectureIds); err != nil {
		log.Println(err)
		resp := studentResp{
			Status:  false,
			Message: "Lectures are not found.",
			Data:    student,
		}
		return c.Status(200).JSON(resp)
	}

	if len(*lectureIds) > 0 {
		err := config.Database.Where(lectureIds).Find(&lectures).Error
		if err != nil {
			log.Println(err)
			resp := studentResp{
				Status:  false,
				Message: "Lectures are not found.",
				Data:    student,
			}
			return c.Status(200).JSON(resp)
		}
	}

	fmt.Println(lectures)
	err := config.Database.Model(&student).Association("Lectures").Replace(lectures)
	if err != nil {
		log.Println(err)
		resp := studentResp{
			Status:  false,
			Message: "An error occurs while student is updated.",
			Data:    student,
		}
		return c.Status(200).JSON(resp)
	}

	resp := studentResp{
		Status:  true,
		Message: "Lectures are updated.",
		Data:    student,
	}
	return c.Status(200).JSON(resp)

}
