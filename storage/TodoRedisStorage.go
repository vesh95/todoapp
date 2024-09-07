package storage

import (
	"container/list"
	"context"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"log"
	"todo/todo"
)

type TodoRedisStorage struct {
	client *redis.Client
	logger *log.Logger
	ctx    context.Context
}

func NewRedisStorage(opt *redis.Options, logger *log.Logger) *TodoRedisStorage {
	rdb := redis.NewClient(opt)

	return &TodoRedisStorage{client: rdb, logger: logger, ctx: context.Background()}
}

// Implementing repository

func (t *TodoRedisStorage) GetAll() *list.List {
	allItems := list.New().Init()
	var cursor uint64
	for {
		keys, cursor, err := t.client.Scan(t.ctx, cursor, "*", 50).Result()
		if err != nil {
			t.logger.Println("GetAll from Redis all keys error: " + err.Error())

			break
		}

		for _, key := range keys {
			val, err := t.client.Get(t.ctx, key).Result()

			if err != nil {
				t.logger.Println("GetAll from Redis get value error: " + err.Error())
				continue
			}

			todoItem := &todo.Todo{}

			err = jsoniter.Unmarshal([]byte(val), todoItem)

			if err != nil {
				t.logger.Println("GetAll from Redis unmarshal error: " + err.Error())
				continue
			}

			allItems.PushBack(todoItem)
		}

		if cursor == 0 {
			break
		}
	}

	return allItems
}

func (t *TodoRedisStorage) Get(uuid uuid.UUID) (*todo.Todo, error) {
	val, err := t.GetByString(uuid.String())

	return val, err
}

func (t *TodoRedisStorage) GetByString(uuid string) (*todo.Todo, error) {
	todoItem := &todo.Todo{}

	val, err := t.client.Get(t.ctx, uuid).Result()
	if err != nil {
		t.logger.Println("Get from Redis error: " + err.Error())

		return nil, err
	}

	err = jsoniter.Unmarshal([]byte(val), todoItem)
	if err != nil {
		t.logger.Println("Get from Redis error: " + err.Error())

		return nil, err
	}

	return todoItem, nil
}

func (t *TodoRedisStorage) Add(todo *todo.Todo) {
	val, err := jsoniter.Marshal(todo)
	if err != nil {
		t.logger.Println("Add to Redis error: " + err.Error())
		return
	}

	_, err = t.client.Set(t.ctx, todo.ID.String(), val, redis.KeepTTL).Result()
	if err != nil {
		t.logger.Println("Add to Redis error: " + err.Error())
		return
	}
}

func (t *TodoRedisStorage) Remove(todo *todo.Todo) {
	// TODO поправить сообщение "redis: nil" на несуществующем ключе
	_, err := t.client.Del(t.ctx, todo.ID.String()).Result()
	if err != nil {
		t.logger.Println("Remove from Redis error: " + err.Error())

		return
	}
}

func (t *TodoRedisStorage) Count() int64 {
	val, err := t.client.DBSize(t.ctx).Result()
	if err != nil {
		t.logger.Println("Count of lines Redis error: " + err.Error())
	}

	return val
}
