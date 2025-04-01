# test_task_bkh_ekom_click_counter
Test task for golang developer


## HOW TO USE
- docker-compose up -d
- go run cmd/migration/main.go
- go run cmd/counter/main.go

## WRK STATS
Running 3m test @ http://localhost:3000/counter/1
  16 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    54.38ms   20.54ms 299.94ms   88.65%
    Req/Sec     1.17k   268.56     2.23k    75.13%
  3347054 requests in 3.00m, 376.66MB read
  Socket errors: connect 0, read 3141, write 0, timeout 0
Requests/sec:  18586.26
Transfer/sec:      2.09MB

## Задача 1. Счетчик кликов.
Есть набор баннеров (от 10 до 100). У каждого есть ИД и название (id, name)

Нужно сделать сервис, который будет считать клики и собирать их в поминутную статистику (timestamp, bannerID, count)

 

Нужно сделать АПИ с двумя методами:

1. /counter/<bannerID> (GET)

Должен посчитать +1 клик по баннеру с заданным ИД

 

2. /stats/<bannerID> (POST)

Должен выдать статистику показов по баннеру за указанный промежуток времени (tsFrom, tsTo)

 

Язык: golang

СУБД: mongo или psql

Сложность:

- junior = кол-во запросов /counter 10-50 в секунду

- middle+ = кол-во запросов /counter 100-500 в секунду

 

PS: тесты делать не обязательно

 

Пример работы сервиса:

Requests logs:

2024-12-12T10:00:10 GET /counter/1

2024-12-12T10:00:20 GET /counter/1

2024-12-12T10:00:30 GET /counter/1

2024-12-12T10:00:40 GET /counter/1

2024-12-12T10:00:50 GET /counter/2

2024-12-12T10:01:00 GET /counter/2

2024-12-12T10:01:10 GET /counter/1

2024-12-12T10:01:11 GET /counter/1

2024-12-12T10:03:22 GET /counter/1

2024-12-12T10:04:33 GET /counter/1

2024-12-12T10:04:44 GET /counter/2

curl -x POST -d '{"from": "2024-12-12T10:00:00", "to": "2024-12-12T10:05:00"}' http://localhost/3000/stats/1

{

"stats": [

{"ts": "2024-12-12T10:00:00", "v": 4},

{"ts": "2024-12-12T10:01:00", "v": 2},

{"ts": "2024-12-12T10:03:00", "v": 1},

{"ts": "2024-12-12T10:04:00", "v": 1},

]

}

curl -x POST -d '{"from": "2024-12-12T10:00:00", "to": "2024-12-12T10:05:00"}' http://localhost/3000/stats/2

{

"stats": [

{"ts": "2024-12-12T10:00:00", "v": 1},

{"ts": "2024-12-12T10:01:00", "v": 1},

{"ts": "2024-12-12T10:04:00", "v": 1},

]

}
