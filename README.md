# TODOAPP
Простое todo приложение на go

## Установка
1. Для работы понадобится докер. https://docs.docker.com/engine/install/
2. Запустить контейнер с redis с пробросом стандартного порта
```
docker run --name redis -d redis -v redis_db:/etc/redis/database -p 6379:6379
```
3. Запустить приложение 
```
docker run --publish=80:8080 --name=todoapp cr.yandex/crp2t1th419fuf0sh6a3/todoapp:1.0
```
4. Функционал работает по пути http://localhost/todo
5. Может принимать следующие запросы
   1. GET /todo
   2. GET /todo/:uuid
   3. POST /todo { "Task": |string| }
   4. DELETE /todo/:uuid
