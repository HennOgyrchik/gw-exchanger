# gRPC Exchanger
## Методы
1. `GetExchangeRates` - получение всех курсов валют
* Возвращает данные курсов валют из БД
2. `GetExchangeRateForCurrency` - получение курса обмена двух любых валют
* Принимает `from_currency, to_currency string`
* Получает курсы валют из БД
* Вычисляет коэффициент обмена `from_currency`/`to_currency`
* Возвращает `from_currency, to_currency string`, `rate float32`

## Конфигурация
Чтение конфигурации происходит из файла, переданного флагом `-c` (по умолчанию - чтение из корня проекта).

Конфигурация подключения к PostgreSQL
*  `PSQL_HOST` - default `localhost`
* `PSQL_PORT` - default `5432`
* `PSQL_DB_NAME` - default `postgres`
* `PSQL_USER` - default `postgres`
* `PSQL_PASSWORD` - default `postgres`
* `PSQL_SSL_MODE` - default `disable`
* `PSQL_CONN_TIMEOUT` - default `60` (в секундах)

Конфигурация gRPC-сервера
* `GRPC_HOST` - default `localhost`
* `GRPC_PORT` - default `9090`
* `GRPC_TIMEOUT` - default `60` (в секундах)