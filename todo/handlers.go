package todo

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoHandler struct {
	todoRepository *TodoRepository
}

func (handler *TodoHandler) GetAll(c *fiber.Ctx) error {
	todos, err := handler.todoRepository.FindAll()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": 404,
			"error":  err.Error(),
		})
	}
	return c.JSON(todos)
}

func (handler *TodoHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	todo, err := handler.todoRepository.FindById(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": 404,
			"error":  err.Error(),
		})
	}
	return c.JSON(todo)
}

func (handler *TodoHandler) Create(c *fiber.Ctx) error {
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Review your payload",
			"error":   err.Error(),
		})
	}
	createdTodo, err := handler.todoRepository.Create(todo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Failed to create todo",
			"error":   err.Error(),
		})
	}
	return c.JSON(createdTodo)
}

func (handler *TodoHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	todoUpdate := Todo{}
	if err := c.BodyParser(&todoUpdate); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Review your payload",
			"error":   err.Error(),
		})
	}
	err := handler.todoRepository.Update(id, &todoUpdate)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "Todo updated",
	})
}

func (handler *TodoHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := handler.todoRepository.Delete(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  err.Error(),
		})
	}
	return c.Status(204).JSON(nil)
}

func NewTodoHandler(repository *TodoRepository) *TodoHandler {
	return &TodoHandler{
		todoRepository: repository,
	}
}

func Register(router fiber.Router, mongo *mongo.Database) {
	todoRepository := NewTodoRepository(mongo)
	todoHandler := NewTodoHandler(todoRepository)

	todoRouter := router.Group("/todo")
	todoRouter.Get("/", todoHandler.GetAll)
	todoRouter.Get("/:id", todoHandler.GetById)
	todoRouter.Post("/", todoHandler.Create)
	todoRouter.Put("/:id", todoHandler.Update)
	todoRouter.Delete("/:id", todoHandler.Delete)
}
