app:
	go build -C cmd/app -o go-ify && ./cmd/app/go-ify

run:
	air

dockerb:
	go tool templ generate
	sudo docker build -t go-ify-image:latest -f ./build/docker/Dockerfile .

dockerp:
	sudo docker image prune -fa

dockerup:
	sudo docker compose -f ./build/docker/docker-compose.yaml up 

dockerdown:
	sudo docker compose -f ./build/docker/docker-compose.yaml down

rebuild:
	go tool templ generate
	sudo docker build -t go-ify-image:latest -f ./build/docker/Dockerfile .
	sudo docker compose -f ./build/docker/docker-compose.yaml up
	sudo docker image prune -fa

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

