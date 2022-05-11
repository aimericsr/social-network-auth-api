postgres:
	docker run -d --name postgres14 -e POSTGRES_USER=rootUser -e POSTGRES_PASSWORD=secretPassword -p 5432:5432 postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=rootUser --owner=rootUser social_network

dropdb:
	docker exec -it postgres14 dropdb festivales

test:
	go test -v -count=1 -cover ./...

image:
	docker build -t auth-api-festivales:latest .

container:
	docker run -d --name auth-api-festivales -p 8080:8080 -e DB_SOURCE="postgresql://root:secretPassword@postgres14:5432/festivales?sslmode=disable" auth-api-festivales:latest

.PHONY: postgres createdb dropdb image container