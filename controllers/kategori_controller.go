package controllers

import (
	"evermos-mini/config"
	"evermos-mini/models"

	"github.com/gofiber/fiber/v2"
)

// GET /kategori
func GetAllKategori(c *fiber.Ctx) error {
	var kategori []models.Category
	config.DB.Find(&kategori)
	return c.JSON(fiber.Map{"status": true, "data": kategori})
}

// POST /kategori
func CreateKategori(c *fiber.Ctx) error {
	var input struct {
		NamaCategory string `json:"nama_category"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

	category := models.Category{NamaCategory: input.NamaCategory}
	config.DB.Create(&category)

	return c.JSON(fiber.Map{"status": true, "message": "Kategori berhasil ditambahkan", "data": category})
}

// PUT /kategori/:id
func UpdateKategori(c *fiber.Ctx) error {
	id := c.Params("id")

	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Kategori tidak ditemukan"})
	}

	var input struct {
		NamaCategory string `json:"nama_category"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Input tidak valid"})
	}

	category.NamaCategory = input.NamaCategory
	config.DB.Save(&category)

	return c.JSON(fiber.Map{"status": true, "message": "Kategori berhasil diperbarui", "data": category})
}

// DELETE /kategori/:id
func DeleteKategori(c *fiber.Ctx) error {
	id := c.Params("id")

	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Kategori tidak ditemukan"})
	}

	config.DB.Delete(&category)
	return c.JSON(fiber.Map{"status": true, "message": "Kategori berhasil dihapus"})
}
