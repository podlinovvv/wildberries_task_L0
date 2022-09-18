## Задание L0


Cервис, а также nats-streaming и база данных postgres запускаются командой 

`docker compose up`

Сервис подписывается на канал nats-streaming и читает данные из него.
Данные хранятся в кэше и дублируются в базу данных postgres.
Через API можно получить данные из кэша по id.

Если запрашиваемые данные хранятся в бд, но отсутствуют в кэше, то кэш будет восстановлен из бд.
В бд postgres в виде json.

API доступен на `localhost:8080`


Скрипт для отправки данных в канал **publisher.go** запускается вручную.

Он отправляет json-данные со случайно сгенерированными полями.

Каждый пятый json отправляется с другой структурой и фильтруется сервисом как невалидный.




