package handlers

import (
	"fmt"
	"net/http"

	"github.com/gastrader/gohouse/ent/user"
	"github.com/gastrader/gohouse/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofiber/fiber/v2"
)

func (r registerRequest) validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Firstname, validation.Required, validation.Length(2,20)),
		validation.Field(&r.LastName, validation.Required, validation.Length(2,20)),
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(6,15)),
	)

}

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

	if err = request.validate(); err != nil{
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "JSON not proper",
		})
	}

	exist, _ := h.Client.User.Query().Where(user.Email(request.Email)).Only(ctx.Context())
	if exist != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "Account exists",
		})
	}

	hashpw, err := utils.HashPassword(request.Password)
	if err != nil{
		return fmt.Errorf("failed to hash user pw: %w", err)
	}

	_, err = h.Client.User.Create().SetEmail(request.Email).SetFirstName(request.Firstname).SetLastName(request.LastName).SetAvatar(request.Avatar).SetPassword(hashpw).Save(ctx.Context())
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