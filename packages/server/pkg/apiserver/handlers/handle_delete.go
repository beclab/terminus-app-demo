package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (h *handler) DeleteNotesPost(ctx *fiber.Ctx) error {
	var deleteId struct {
		ID int `json:"id"`
	}

	err := ctx.BodyParser(&deleteId)
	if err != nil {
		log.Error("parse delete note error, ", err)
		return err
	}

	err = h.db.Delete(deleteId.ID)
	if err != nil {
		log.Error("delete note error, ", err)

		return err
	}

	return ctx.JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "Delete note success",
	})
}
