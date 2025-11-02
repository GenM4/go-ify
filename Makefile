app:
	go build -C cmd/app -o go-ify && ./cmd/app/go-ify

run:
	air

startdb:
	sudo service postgresql start

pgshell:
	sudo -u postgres psql

gooseup:
	cd ./sql/schema && goose postgres postgres://postgres:postgres@localhost:5432/hlhp up && cd ../..

goosedown:
	cd ./sql/schema && goose postgres postgres://postgres:postgres@localhost:5432/hlhp down && cd ../..

sqlc:
	sqlc generate

dockerb:
	docker build -t go-ify-image:latest -f ./build/docker/Dockerfile .
	docker image prune -fa

dockerup:
	docker compose -f ./build/docker/docker-compose.yaml up

dockerdown:
	docker compose -f ./build/docker/docker-compose.yaml down

rebuild:
	sudo docker build -t go-ify-image:latest -f ./build/docker/Dockerfile .
	sudo docker compose -f ./build/docker/docker-compose.yaml up
	sudo docker image prune -fa
