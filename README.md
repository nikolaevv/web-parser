# Парсер вебсайтов на Go
Скрипт проходит по слайсу ссылок, затем получает тело ответа в GET запросе и ищет в нём количество вхождений подстроки "Go".
Это происходит параллельно сразу для нескольких запросов, но не более, чем для k запросов.

- Для синхронизации горутин используется Waitgroup
- Для ограничения максимального количества параллельных запросов используется буфферезированный канал (блокируется при достижении максимального кол-ва горутин)
- Результат сохраняется в потокобезопасную Map (реализована с Mutex)
