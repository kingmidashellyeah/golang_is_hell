# Простой веб-сервис для вычисления арифметических выражений  


У сервиса 1 endpoint с url-ом /api/v1/calculate. Пользователь отправляет на этот url запрос с телом:

    {
        "expression": "выражение, которое ввёл пользователь"
    }

В ответ пользователь получает HTTP-ответ с телом:

    {
        "result": "результат выражения"
    }

и кодом [200], если выражение вычислено успешно.  

## Структура проекта:

cmd/calc_service - директория с кодом сервера main.go  
internal/calculator - директория с кодом калькулятора  

## Запуск  

Установите Golang https://go.dev/dl/  
Установите Git https://git-scm.com/downloads  
C помощью командной строки клонируйте проект с GitHub:  
git clone https://github.com/kingmidashellyeah/golang_is_hell
Перейдите в директорию calculator_service с проектом и запустите сервер с помощью команды:  
go run ./cmd/calc_service/  

## Работа с сервисом  

Для работы с данным сервисом используйте командную строку.  
Для корректной работы на Windows необходимо использовать Git Bash (устанавливается вместе с Git).  

Пример запроса (вместо "..." нужно вставить выражение):

    curl --location 'http://127.0.0.1:8080/api/v1/calculate' \
    --header 'Content-Type: application/json' \
    --data '{
        "expression":"..."
    }'

Рассмотрим примеры запросов  

#### Корректный запрос:

Введя данный запрос:

    curl --location 'http://127.0.0.1:8080/api/v1/calculate' \
    --header 'Content-Type: application/json' \
    --data '{
        "expression":"1 + 1 * 4"
    }'

вы получите ответ:

    {"result":5}

с кодом [200].

#### Запрос с методом не POST:

Введя данный запрос:

    curl --location --request GET 'http://127.0.0.1:8080/api/v1/calculate' \
    --header 'Content-Type: application/json' \
    --data '{
        "expression":"1 + 1"
    }'

вы получите ответ:

    {"error": "Invalid request method"}

с кодом [405].

#### Запрос с неправильным телом:

Введя данный запрос:

    curl --location 'http://127.0.0.1:8080/api/v1/calculate' \
    --header 'Content-Type: application/json' \
    --data '{
        "expression":"1 + 1
    }'

вы получите ответ:

    {"error": "Invalid request body"}

с кодом [400].

#### Запрос с делением на 0:

Введя данный запрос:

    curl --location 'http://127.0.0.1:8080/api/v1/calculate' \
    --header 'Content-Type: application/json' \
    --data '{
        "expression":"1 / 0"
    }'

вы получите ответ:

    {"error":"division by zero"}

с кодом [422].

#### Запрос с выражением с буквами:

Введя данный запрос:

    curl --location 'http://127.0.0.1:8080/api/v1/calculate' \
    --header 'Content-Type: application/json' \
    --data '{
        "expression":"1 / a"
    }'

вы получите ответ:

    {"error": "invalid expression"}

с кодом [422].

#### иные ошибки:

В случае иной ошибки на стороне сервера будет получен ответ:

    {"error":"Internal server error"}  
с кодом [500].
