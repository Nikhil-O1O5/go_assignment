package handler

import (
	"context"
	"go-backend-task/db/sqlc"
	"go-backend-task/internal/logger"
	"go-backend-task/internal/models"
	"go-backend-task/internal/repository"
	"go-backend-task/internal/service"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserHandler struct {
	Repo      *repository.UserRepository
	Validator *validator.Validate
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		Repo:      repo,
		Validator: validator.New(),
	}
}

// CreateUser handles POST /users [cite: 31]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.UserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.Validator.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid date format. Use YYYY-MM-DD"})
	}

	user, err := h.Repo.Create(context.Background(), db.CreateUserParams{
		Name: req.Name,
		Dob:  dob,
	})
	if err != nil {
		logger.Log.Error("Failed to create user", zap.Error(err))
		return c.Status(500).JSON(fiber.Map{"error": "Could not create user"})
	}

	age, _ := service.CalculateAge(req.DOB)

	return c.Status(201).JSON(models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  age,
	})
}

// GetUser handles GET /users/:id [cite: 44]
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	user, err := h.Repo.GetByID(context.Background(), int32(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	dobStr := user.Dob.Format("2006-01-02")
	age, _ := service.CalculateAge(dobStr)

	return c.JSON(models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  dobStr,
		Age:  age,
	})
}

// UpdateUser handles PUT /users/:id [cite: 53]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var req models.UserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	dob, _ := time.Parse("2006-01-02", req.DOB)

	user, err := h.Repo.Update(context.Background(), db.UpdateUserParams{
		ID:   int32(id),
		Name: req.Name,
		Dob:  dob,
	})
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
	})
}

// DeleteUser handles DELETE /users/:id [cite: 66]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := h.Repo.Delete(context.Background(), int32(id)); err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.SendStatus(fiber.StatusNoContent) // HTTP 204 [cite: 68]
}

// ListUsers handles GET /users [cite: 70]
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.Repo.ListAll(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch users"})
	}

	var response []models.UserResponse
	for _, u := range users {
		dobStr := u.Dob.Format("2006-01-02")
		age, _ := service.CalculateAge(dobStr)
		response = append(response, models.UserResponse{
			ID:   u.ID,
			Name: u.Name,
			DOB:  dobStr,
			Age:  age,
		})
	}
	return c.JSON(response)
}