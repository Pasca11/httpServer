.PHONY: build run docker-build docker-run clean test

# Переменные
APP_NAME=water-service
DOCKER_IMAGE=$(APP_NAME):latest

# Сборка приложения локально
build:
	go build -o $(APP_NAME) main.go

# Запуск приложения локально
run:
	go run main.go

# Сборка Docker образа
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Запуск Docker контейнера
docker-run:
	docker run -p 8080:8080 $(DOCKER_IMAGE)

# Очистка бинарных файлов
clean:
	rm -f $(APP_NAME)
	go clean

# Запуск тестов
test:
	go test -v ./...

# Установка зависимостей
deps:
	go mod download

# Проверка кода
lint:
	go vet ./...
	
# Помощь
help:
	@echo "Доступные команды:"
	@echo "  make build         - Собрать приложение локально"
	@echo "  make run          - Запустить приложение локально"
	@echo "  make docker-build - Собрать Docker образ"
	@echo "  make docker-run   - Запустить Docker контейнер"
	@echo "  make clean        - Очистить бинарные файлы"
	@echo "  make test         - Запустить тесты"
	@echo "  make deps         - Установить зависимости"
	@echo "  make lint         - Проверить код" 

docker-build:
	docker build -t http_server:latest .

docker-run:
	docker run -d -p 8082:8082 http_server