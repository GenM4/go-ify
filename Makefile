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

dockerc:
	docker build -t homelab-homepage-image .
	docker stop hhl
	docker rm hhl
	docker container create -p 8080:8080 --name hhl homelab-homepage-image

dockerup:
	docker start hhl

dockerdown:
	docker stop hhl
