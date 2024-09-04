package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"todo/todo"
)

type TodoHandler struct {
	Repository todo.Repository
}

type TodoRequest struct {
	Task string `json:"task"`
}

type TodoListResponse struct {
	Data []*todo.Todo `json:"data"`
}

func (t *TodoHandler) Index(c *gin.Context) {
	itemsList := t.Repository.GetAll()
	var data = make([]*todo.Todo, itemsList.Len())
	i := 0
	for e := itemsList.Front(); e != nil; e = e.Next() {
		data[i] = e.Value.(*todo.Todo)
		i++
	}

	c.JSON(http.StatusOK, TodoListResponse{Data: data})
}

func (t *TodoHandler) Store(c *gin.Context) {
	var todoRequest *TodoRequest
	if err := c.ShouldBindJSON(&todoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	newItem := &todo.Todo{ID: uuid.New(), Task: todoRequest.Task, IsComplete: false}
	t.Repository.Add(newItem)

	c.JSON(http.StatusCreated, newItem)

	return
}

func (t *TodoHandler) Show(c *gin.Context) {
	id := c.Param("id")

	if todoItem, err := t.Repository.GetByString(id); err == nil && todoItem != nil {
		c.JSON(http.StatusOK, todoItem)
	} else {
		c.JSON(http.StatusNotFound, struct {
			Error string `json:"error"`
		}{Error: err.Error()})
	}
}

func (t *TodoHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if todoItem, err := t.Repository.GetByString(id); err == nil && todoItem != nil {
		t.Repository.Remove(todoItem)
		c.JSON(http.StatusOK, struct {
		}{})
	} else {
		c.JSON(http.StatusNotFound, struct {
			Error string `json:"error"`
		}{Error: err.Error()})
	}

}
