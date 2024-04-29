package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h *handler) ListNotesGet(ctx *fiber.Ctx) error {
	list, err := h.db.List()
	if err != nil {
		return err
	}

	return ctx.JSON(list)
}
