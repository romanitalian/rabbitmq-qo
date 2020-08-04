# rabbitmq-go

Пример работы с rabbitmq:
- коннект (dial)
- публикация (publisher)
- получения сообщения (consumer)


#### HOW TO:


1. Запустить сервер rabbitmq
> docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management

2. Опубликовать сообщения (зашито в коде)
> go run cmd/publisher/main.go

3. Прочитать сообщение из rabbitmq
> go run cmd/consumer/main.go
