package noteHandler

import (
	"hungdim2001/database"
	"hungdim2001/internal/model"
	"hungdim2001/validate"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetNotes(c * fiber.Ctx) error {
    db:= database.DB
    var notes[]model.Note
    db.Find( & notes)
    if len(notes) == 0 {
        return c.Status(404).JSON(fiber.Map {
            "status": "error",
            "message": "No notes present",
            "data": nil,
        })
    }
    return c.JSON(fiber.Map {
        "status": "success",
        "message": "Notes Founds",
        "data": notes,
    })
}
func CreateNotes(c * fiber.Ctx) error {
    db:= database.DB
    note:= new(model.Note)
    noteFound:= new(model.Note)
    err:= c.BodyParser(note)
    note.ID = uuid.New()

    errors:= validate.ValidateStruct( * note)
    if errors != nil {
        return c.Status(fiber.StatusBadRequest).JSON(errors)
    }
// check existing
    db.Find( & noteFound, "title = ?", note.Title)
    if (noteFound.Title == note.Title) {
        return c.Status(500).JSON(fiber.Map {
            "status": "error",
            "message": "Title is exist",
            "data": nil,
        })
    } else {
        err = db.Create( & note).Error
        if err != nil {
            return c.Status(500).JSON(fiber.Map {
                "status": "error",
                "message": "Can not create note",
                "data": err,
            })
        }

        return c.JSON(fiber.Map {
            "status": "success",
            "message": "Created Note",
            "data": note,
        })
    }
}
func GetNote(c * fiber.Ctx) error {
    db:= database.DB
    var note model.Note

    // Read the param noteId
    id:= c.Params("noteId")

    // Find the note with the given Id
        db.Find( & note, "id = ?", id)

    // If no such note present return an error
    if note.ID == uuid.Nil {
        return c.Status(404).JSON(fiber.Map {
            "status": "error",
            "message": "No note present",
            "data": nil,
        })
    }

    // Return the note with the Id
    return c.JSON(fiber.Map {
        "status": "success",
        "message": "Notes Found",
        "data": note,
    })
}
func UpdateNote(c * fiber.Ctx) error {
    type updateNote struct {
        Title string `json:"Title"`
        SubTitle string `json:"Sub_title"`
        Text string `json:"Text"`
    }
    db:= database.DB
    var note model.Note

    // Read the param noteId
    id:= c.Params("noteId")

    // Find the note with the given Id
    db.Find( & note, "id = ?", id)

    // If no such note present return an error
    if note.ID == uuid.Nil {
        return c.Status(404).JSON(fiber.Map {
            "status": "error",
            "message": "No note present",
            "data": nil,
        })
    }

    // Store the body containing the updated data and return error if encountered
    var updateNoteData updateNote
    err:= c.BodyParser( & updateNoteData)
    if err != nil {
        return c.Status(500).JSON(fiber.Map {
            "status": "error",
            "message": "Review your input",
            "data": err,
        })
    }

    // Edit the note
    note.Title = updateNoteData.Title
    note.SubTitle = updateNoteData.SubTitle
    note.Text = updateNoteData.Text

    // Save the Changes
    db.Save( &note)

    // Return the updated note
    return c.JSON(fiber.Map {
        "status": "success",
        "message": "Notes Found",
        "data": note,
    })
}

func DeleteNote(c * fiber.Ctx) error {
    db:= database.DB
    var note model.Note

    // Read the param noteId
    id:= c.Params("noteId")

    // Find the note with the given Id
        db.Find( & note, "id = ?", id)

    // If no such note present return an error
    if note.ID == uuid.Nil {
        return c.Status(404).JSON(fiber.Map {
            "status": "error",
            "message": "No note present",
            "data": nil,
        })
    }

    // Delete the note and return error if encountered
    err:= db.Delete( & note, "id = ?", id).Error

        if err != nil {
        return c.Status(404).JSON(fiber.Map {
            "status": "error",
            "message": "Failed to delete note",
            "data": nil,
        })
    }

    // Return success message
    return c.JSON(fiber.Map {
        "status": "success",
        "message": "Deleted Note",
    })
}