package handlers

import (
	"net/http"

	"github.com/gastrader/gohouse/ent/user"
	"github.com/gastrader/gohouse/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) UserRegister(ctx *fiber.Ctx) error {
	var request registerRequest
	err := ctx.BodyParser(&request); if err != nil{
		err = ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "invalid JSON",
		}); if err != nil{
			ctx.Status(http.StatusInternalServerError)
		}
		
	}

	exist, _ := h.Client.User.Query().Where(user.Email(request.Email)).Only(ctx.Context())
	if exist != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "Account exists",
		})
	}

	_, err = h.Client.User.Create().SetEmail(request.Email).SetFirstName(request.Firstname).SetLastName(request.LastName).SetAvatar(request.Avatar).SetPassword(request.Password).Save(ctx.Context())
	if err != nil{
		utils.Errorf("fail to creat user: %v", err)
		return nil
	}

	_ = ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"error": false,
		"message": "Registration successful",
	})
	return nil
}