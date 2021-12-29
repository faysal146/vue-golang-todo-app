package main

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// model
type Todo struct {
	Id        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// database
var todos []Todo = []Todo{}

func main() {

	app := fiber.New()
	app.Use(cors.New())

	apiV1 := app.Group("/api/v1")

	apiV1.Get("/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	apiV1.Post("/todos", func(c *fiber.Ctx) error {
		var newTodo Todo
		newTodo.CreatedAt = time.Now()
		newTodo.UpdatedAt = time.Now()
		newTodo.Id = len(todos) + 1
		json.Unmarshal(c.Body(), &newTodo)
		todos = append(todos, newTodo)
		return c.JSON(newTodo)
	})

	apiV1.Patch("/todos/:id", func(c *fiber.Ctx) error {
		todoId := c.Params("id")
		num, err := strconv.ParseInt(todoId, 32, 0)
		if err != nil {
			panic(err)
		}
		var updatedTodo = todos[num-1]
		updatedTodo.UpdatedAt = time.Now()
		json.Unmarshal(c.Body(), &updatedTodo)

		todos[num-1] = updatedTodo
		return c.JSON(updatedTodo)
	})

	apiV1.Delete("/todos/:id", func(c *fiber.Ctx) error {
		todoId := c.Params("id")
		num, err := strconv.ParseInt(todoId, 32, 0)
		if err != nil {
			panic(err)
		}

		var tt []Todo
		var toRemoveTodo Todo
		for _, val := range todos {
			if int64(val.Id) != num {
				tt = append(tt, val)
			} else {
				toRemoveTodo = val
			}
		}
		todos = tt
		return c.JSON(toRemoveTodo)
	})

	app.Listen("127.0.0.1:8000")
}
