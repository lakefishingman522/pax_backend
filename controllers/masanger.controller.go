package controllers

import (
	"fmt"
	"hyperpage/utils"

	"github.com/gofiber/fiber/v2"
)

// MakeCall обрабатывает запрос на создание звонка
func MakeCall(c *fiber.Ctx) error {

	// Извлекаем данные о вызове из тела запроса
	var callData struct {
		Token   string `json:"token"`
		Payload string `json:"payload"`
	}

	if err := c.BodyParser(&callData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный формат данных",
		})
	}

	err := utils.VoipCall(callData.Token, callData.Payload)
	if err != nil {
		fmt.Println("Error:", err)
		// Handle the error as needed
	}

	response := map[string]string{"message": "Звонок успешно создан"}

	// Отправляем ответ в формате JSON
	return c.JSON(response)
}

func StopCall(c *fiber.Ctx) error {

	// Извлекаем данные о вызове из тела запроса
	var callData struct {
		Token   string `json:"token"`
		Payload string `json:"payload"`
	}

	if err := c.BodyParser(&callData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный формат данных",
		})
	}

	err := utils.VoipCall(callData.Token, callData.Payload)
	if err != nil {
		fmt.Println("Error:", err)
		// Handle the error as needed
	}

	response := map[string]string{"message": "Звонок успешно завершен"}

	// Отправляем ответ в формате JSON
	return c.JSON(response)
}
