build:
	docker compose build

run:
	docker compose up

docker:
	docker compose run --service-ports web bash