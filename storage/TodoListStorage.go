package storage

import (
	"container/list"
	"errors"
	"github.com/google/uuid"
	"todo/todo"
)

type TodoListStorage struct {
	data *list.List
}

func NewTodoListStorage() *TodoListStorage {
	return &TodoListStorage{data: list.New().Init()}
}

func (t *TodoListStorage) GetAll() *list.List {
	return t.data
}

func (t *TodoListStorage) Get(uuid uuid.UUID) (*todo.Todo, error) {
	for e := t.data.Front(); e != nil; e.Next() {
		currentTodo := e.Value.(*todo.Todo)
		if currentTodo.ID == uuid {
			return currentTodo, nil
		}
	}

	return nil, errors.New("not found")
}

func (t *TodoListStorage) Add(todo *todo.Todo) {
	t.data.PushBack(todo)
}

func (t *TodoListStorage) Remove(todo *todo.Todo) {
	for e := t.data.Front(); e != nil; e.Next() {
		if e.Value == todo {
			t.data.Remove(e)
			return
		}
	}
}

func (t *TodoListStorage) Count() int {
	return t.data.Len()
}

func (t *TodoListStorage) GetByString(uuid string) (*todo.Todo, error) {
	for e := t.data.Front(); e != nil; e = e.Next() {
		currentTodo := e.Value.(*todo.Todo)
		if currentTodo.ID.String() == uuid {
			return currentTodo, nil
		}
	}

	return nil, errors.New("not found")
}
