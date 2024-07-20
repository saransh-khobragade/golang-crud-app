package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos map[string]Todo
var mutex sync.Mutex

func main() {
	todos = make(map[string]Todo)
	router := gin.Default()

	// Create a new Todo (POST /todos)
	router.POST("/todos", createTodo)

	// Get all Todos (GET /todos)
	router.GET("/todos", getAllTodos)

	// Get a specific Todo by ID (GET /todos/:id)
	router.GET("/todos/:id", getTodoByID)

	// Update a Todo by ID (PUT /todos/:id)
	router.PUT("/todos/:id", updateTodo)

	// Delete a Todo by ID (DELETE /todos/:id)
	router.DELETE("/todos/:id", deleteTodo)

	router.Run(":8080")
}

func createTodo(c *gin.Context) {
	var newTodo Todo

	// Parse request body into newTodo object
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a unique ID
	mutex.Lock()
	newTodo.ID = generateID()
	mutex.Unlock()

	todos[newTodo.ID] = newTodo

	c.JSON(http.StatusCreated, newTodo)
}

func getAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func getTodoByID(c *gin.Context) {
	id := c.Param("id")
	todo, found := todos[id]

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")
	_, found := todos[id]

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// Parse request body into updated todo data
	var updatedTodo Todo
	if err := c.BindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTodo.ID = id // Keep original ID

	todos[id] = updatedTodo

	c.JSON(http.StatusOK, updatedTodo)
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	_, found := todos[id]

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	delete(todos, id)

	c.JSON(http.StatusNoContent, nil)
}

func generateID() string {
	// Replace this with a more robust ID generation mechanism
	return "TODO: Implement unique ID generation"
}
