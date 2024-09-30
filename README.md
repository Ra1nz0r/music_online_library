<div align="center"> <h1 align="center"> ТЗ: Реализация онлайн библиотеки песен 🎶 </h1> </div>

__Представляет собой онлайн библиотеку песен, где пользователи могут просматривать тексты песен различных исполнителей.__

- __В проекте реализованы REST методы__:
    - [x] Добавление песни[^1].
    - [x] Получение данных библиотеки с фильтрацией по всем полям и пагинацией.
    - [x] Получение текста песни с пагинацией по куплетам[^2].
    - [x] Удаление песни.
    - [x] Изменение параметров песни.

[Инструкция по локальному запуску и информация по приложению.](#local)\
[Инструкция по созданию Docker образа и запуску контейнера.](#docker)\
[Инструкция по запуску PostgreSQL.](#postgresql)

***
#### Инструкция по локальному запуску и информация по приложению.

_Для изменения стандартных параметров, нужно изменить значения в ```.env``` файле корня проекта._
</div>

По-умолчанию приложение запускается на ```localhost:7654```

- Программу можно запускать двумя способами через терминал.
    - Обычные команды. 
    - Короткими командами из TaskFile.
<div>

- ___Для запуска приложения в терминале.___\
```go run ./cmd/app``` или ```task run```
<div>

- ___Для запуска тестов в терминале.___\
```go test -v ./... -count=1``` или ```task test```

***
[^1]: При добавлении песни, происходит подключение ко внешнему API для получения дополнительных данных. Если запрос завершается неудачей, то песня будет добавлена без дополнительных параметров.

[^2]: Текст разбивается на куплеты по символу '\n\n', в самих же куплетах символ '\n' заменяется переносом на новую строчку.