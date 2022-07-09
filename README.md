# SCHEDULIZER

Приложение, которое реализует выполнение задач в параллельном и последовательном режиме по правилам, заданным во внешнем файле или запросами к web серверу. 
## 40
Напишите свой веб-сервер, принимающий на вход задачу, но обрабатывающий параллельно поданные задачи последовательно, как будто на одном процессоре.

add-метод POST:

Добавляет новую задачу с длительностью timeDuration в список задач, с параметром sync/async:

sync — держим HTTP-коннект и возвращаем ответ только после выполнения всех задач в очереди, а также выполнения самой задачи, которая держит коннект;
async POST — добавляем задачу в очередь и сразу отключаемся.
schedule-метод GET возвращает массив актуальных задач, стоящих в очереди на выполнение, в формате JSON.

time-метод GET возвращает оставшееся время на выполнение всех находящихся в очереди задач.

Помимо веб-сервера, напишите скрипт на Go, тестирующий ваш веб-сервер и проверяющий все основные юзкейсы. (Нужно в несколько горутин ддосить ваш веб-сервер post-запросами и проверять, совпадает ли таймер с теоретическим временем, рассчитанным на клиенте.)

## Сборка и запуск приложения
Склонировать репозиторий, запустить сборку и собранное приложение
```bash
go clean && go build && ./schedulizer
``` 
## Проверка работы
### Добавление синхронной задачи
```bash
curl --location --request POST 'http://localhost:8090/add' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'mode=sync' \
--data-urlencode 'duration=5s'
```
Придет ответ вида
```bash
{"duration":5000000000,"human_duration":"5s","Mode":"sync"}
```
Ответ от сервера придет только после завершения задачи.

В консоли приложения будет выведено
```bash
2022/07/10 00:23:03 Start task #1 duration: 5s
2022/07/10 00:23:08 Stop task  #1 after 5s
```

### Добавление aсинхронной задачи
```bash
curl --location --request POST 'http://localhost:8090/add' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'mode=аsync' \
--data-urlencode 'duration=5s'
```
Придет ответ вида
```bash
{"duration":5000000000,"human_duration":"5s","Mode":"async"}
```
Ответ от сервера придет сразу после постановки задачи в очередь.

В консоли приложения будет выведено
```bash
2022/07/10 00:23:03 Start task #2 duration: 5s
2022/07/10 00:23:08 Stop task  #2 after 5s
```

### Просмотр очереди задач и суммарного времени их выполнения
Запустить несколько асинхронных задач большой длительности
```bash
curl --location --request POST 'http://localhost:8090/add' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'mode=async' \
--data-urlencode 'duration=15s'

curl --location --request POST 'http://localhost:8090/add' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'mode=async' \
--data-urlencode 'duration=15s'

curl --location --request POST 'http://localhost:8090/add' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'mode=async' \
--data-urlencode 'duration=15s'
```
В консоли приложения будет выведено
```bash
2022/07/10 00:27:44 Start task #1 duration: 15s
2022/07/10 00:27:45 Queued 2 tasks. Current run schedulled
2022/07/10 00:27:46 Queued 3 tasks. Current run schedulled
```
Выполнить запрос 
```bash
curl --location --request GET 'http://localhost:8090/schedule'
```
В ответ вернется список из поставленных в очередь задач 
```bash
[{"duration":15000000000,"human_duration":"15s","Mode":"async"},{"duration":15000000000,"human_duration":"15s","Mode":"async"},{"duration":15000000000,"human_duration":"15s","Mode":"async"}]%                    
```
Первая выполняемая в списке будет отсутствовать

Выполнить запрос 
```bash
curl --location --request GET 'http://localhost:8090/time'
```
В ответ вернется суммарная длительность задач в очереди
```bash
{"Durations":45000000000,"HumanDurations":"45s"}
```
Длительность выполняемой задачи не включается в сумму.

## Тестирование
Описанный выше алгоритм проверки реализуется в тесте `TestAsyncAddTask`. Запустить тесты можно командой 
```bash
go test
```
Вывод в консоли теста
```bash
PASS
ok      schedulizer     18.141s
```
Вывод в консоли приложения
```bash
2022/07/10 00:40:51 Start task #1 duration: 1s
2022/07/10 00:40:51 Queued 2 tasks. Current run schedulled
2022/07/10 00:40:52 Stop task  #1 after 1s
2022/07/10 00:40:52 Start task #2 duration: 2s
2022/07/10 00:40:54 Stop task  #2 after 2s
2022/07/10 00:40:54 Start task #3 duration: 3s
2022/07/10 00:40:54 Queued 4 tasks. Current run schedulled
2022/07/10 00:40:54 Queued 4 tasks. Current run schedulled
2022/07/10 00:40:54 Queued 4 tasks. Current run schedulled
2022/07/10 00:40:54 Queued 5 tasks. Current run schedulled
2022/07/10 00:40:54 Queued 6 tasks. Current run schedulled
2022/07/10 00:40:57 Stop task  #3 after 3s
2022/07/10 00:40:57 Start task #4 duration: 3s
2022/07/10 00:41:00 Stop task  #4 after 3s
2022/07/10 00:41:00 Start task #5 duration: 3s
2022/07/10 00:41:03 Stop task  #5 after 3s
2022/07/10 00:41:03 Start task #6 duration: 3s
2022/07/10 00:41:06 Stop task  #6 after 3s
2022/07/10 00:41:06 Start task #7 duration: 3s
2022/07/10 00:41:09 Stop task  #7 after 3s
2022/07/10 00:41:09 Start task #8 duration: 0s
2022/07/10 00:41:09 Stop task  #8 after 0s
```