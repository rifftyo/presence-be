package handler

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rifftyo/presence-be/internal/delivery/http/request"
	"github.com/rifftyo/presence-be/internal/delivery/http/response"
	"github.com/rifftyo/presence-be/internal/usecase"
	"github.com/rifftyo/presence-be/utils"
)

type UserHandler struct {
	usecase usecase.UserUseCase
}

func NewUserHandler(usecase usecase.UserUseCase) *UserHandler {
	return &UserHandler{usecase}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req request.RegisterUserRequest

	req.Name = c.FormValue("name")
	req.Email = c.FormValue("email")
	req.Password = c.FormValue("password")
	req.Telephone = c.FormValue("telephone")
	req.RoleId = c.FormValue("role_id")

	fileHeader, err := c.FormFile("image_profile")
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		defer file.Close()

		fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), fileHeader.Filename)

		filePath, err := utils.SaveFile(file, fileName)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		req.ImageProfile = filePath
	}
	
	user, token, err := h.usecase.Register(&req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	resp := response.UserResponse{
		Message: "User registered successfully",
		Token: token,
		User: utils.MapUserToUserResponse(user),
	}

	return c.JSON(resp)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req request.LoginUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	user, token, err := h.usecase.Login(&req)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	resp := response.UserResponse{
		Message: "User login successfully",
		Token: token,
		User: utils.MapUserToUserResponse(user),
	}

	return c.JSON(resp)
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
    userID := c.Locals("userID").(string)

    user, err := h.usecase.GetUserByID(userID)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "user not found",
        })
    }

    resp := utils.MapUserToUserDetailResponse(user)

    return c.Status(fiber.StatusOK).JSON(resp)
}