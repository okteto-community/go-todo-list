.PHONY: start
start:
	go run main.go

.PHONY: debug
debug:
	dlv debug --headless --listen=:2345 --log --api-version=2

.PHONY: build
build:
	go build -o app

.PHONY: docker-up
docker-up:
	docker-compose up --build

.PHONY: docker-down
docker-down:
	docker-compose down

.PHONY: docker-logs
docker-logs:
	docker-compose logs -f

.PHONY: test-api
test-api:
	./test-api.sh

.PHONY: okteto-deploy
okteto-deploy:
	okteto deploy --wait

.PHONY: okteto-destroy
okteto-destroy:
	okteto destroy

.PHONY: clean
clean:
	rm -f app
	docker-compose down -v

