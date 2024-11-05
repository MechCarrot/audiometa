# audimeta CLI
CLI tool to extract metadata from audio files.


# Структура проекта
В проекте используется domain-driven design  
cmd - entrypoint (api + cli)  
extractors - пакеты, связанные с извлечением метаданных из аудио файлов  
models - domain сущности (аудио, метаданные)  
services - сервисы по извлечению данных, валидации метаданных, event-listener сервер  
storage - интерфейсы для хранения данных в хранилище (работа с бд)  


# Usage
## API:  

Для запуска API server'a следует выполнить следующую команду:
```bash
go run cmd/api/main.go
``` 
Существует 3 endpoint'а для API:  
[ 1 ] GET /list - Получить список всех доступных аудиофайлов  
[ 2 ] GET /get/{id} - Получить метаданные по id  
[ 3 ] POST /upload - Загрузить файл с метаданными  

## CLI:

Entrypoint - cmd/cli/main.go  
Структура CLI:  
./application <get|update|list> [-id=.. | -filename=.. | ]  


# TODO
[  ] Тесты (mock,unit)
[  ] Docker
[  ] Использование Cobra/Viper
[  ] Улучшение storage
