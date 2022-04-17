package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/nuveo-test/entity"
	"github.com/rafaelbreno/nuveo-test/internal"
	"github.com/rafaelbreno/nuveo-test/queue"
	"github.com/rafaelbreno/nuveo-test/services/api/repository"
	"github.com/rafaelbreno/nuveo-test/services/api/storage"
)

type (
	// UserHandler handles all http requests
	// related to the User entity.
	UserHandler struct {
		repo repository.UserRepoI
		in   *internal.Internal
	}
)

// NewUserHandler receives all values necessary
// to build a handler
func NewUserHandler(st *storage.Storage, in *internal.Internal, queue *queue.Queue) UserHandler {
	return UserHandler{
		repo: repository.NewUserRepo(st, in, queue),
		in:   in,
	}
}

// Create handles POST request,
// receives a JSON and insert it in DB.
func (uh *UserHandler) Create(c *fiber.Ctx) error {
	u := new(entity.User)

	if err := json.Unmarshal(c.Body(), u); err != nil {
		uh.in.L.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	createdUser, err := uh.repo.Create(*u)

	if err != nil {
		uh.in.L.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.
		Status(http.StatusCreated).
		JSON(createdUser)
}

// update handles put/patch request,
// receives a json with updated fields
// of a row in db.
func (uh *UserHandler) Update(c *fiber.Ctx) error {
	u := new(entity.User)

	id := c.Params("uuid")

	if err := json.Unmarshal(c.Body(), u); err != nil {
		uh.in.L.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	updatedUser, err := uh.repo.Update(id, *u)

	if err != nil {
		uh.in.L.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.
		Status(http.StatusOK).
		JSON(updatedUser)
}

// Read handles a request with GET type,
// and retrieve the row with specific uuid.
func (uh *UserHandler) Read(c *fiber.Ctx) error {
	id := c.Params("uuid")

	user, err := uh.repo.Read(id)

	if err != nil {
		uh.in.L.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.
		Status(http.StatusOK).
		JSON(user)
}

// ReadAll handles a request with GET type,
// and retrieve all rows from DB.
func (uh *UserHandler) ReadAll(c *fiber.Ctx) error {
	users, err := uh.repo.ReadAll()

	if err != nil {
		uh.in.L.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.
		Status(http.StatusOK).
		JSON(users)
}

// Delete handles a request with DELETE type,
// receives a uuid and removes the given item
// from DB.
func (uh *UserHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("uuid")

	err := uh.repo.Delete(id)

	if err != nil {
		uh.in.L.Error(err.Error())
		return c.
			Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.
		Status(http.StatusOK).
		JSON(fiber.Map{
			"message": fmt.Sprintf("user with id %s deleted successfully", id),
		})
}
