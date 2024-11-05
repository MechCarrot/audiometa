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

CLI: get, upload, delete