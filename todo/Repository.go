package todo

import (
	"container/list"
	"github.com/google/uuid"
)

type Repository interface {
	GetAll() *list.List
	Get(uuid uuid.UUID) (*Todo, error)
	GetByString(uuid string) (*Todo, error)
	Add(todo *Todo)
	Remove(todo *Todo)
	Count() int64
}
