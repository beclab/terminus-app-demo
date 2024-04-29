package handlers

import (
	"github.com/beclab/terminus-app-demo/pkg/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (h *handler) CreateNotesPost(ctx *fiber.Ctx) error {
	var note model.Note
	err := ctx.BodyParser(&note)
	if err != nil {
		log.Error("parse create note error, ", err)
		return err
	}

	newNote, err := h.db.Create(&note)
	if err != nil {
		log.Error("create note error, ", err)
		return err
	}

	return ctx.JSON(newNote)
}
