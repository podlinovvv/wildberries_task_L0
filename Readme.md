### Задание L0

Cервис, а также nats-streaming и база данных postgres запускаются командой 


`docker compose up`


Скрипт для отправки данных в канал **publisher.go** запускается вручную.


Он отправляет json-данные со случайно сгенерированными полями.


Каждый пятый json отправляется с другой структурой и фильтруется сервисом как невалидный.