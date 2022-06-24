postgres:
	docker run -d --name postgres14 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root festivals

dropdb:
	docker exec -it postgres14 dropdb festivals

test:
	go test -v -count=1 -cover ./...

image:
	docker build -t eu.gcr.io/kubernetessandbox-352207/festivals-api:0.0.2 --platform linux/amd64 .

push:
	docker push eu.gcr.io/kubernetessandbox-352207/festivals-api:0.0.2

container:
	docker run --name gcp-api-festivales -p 8080:8080 -e DB_SOURCE="postgresql://postgres:kygCew-porwyt-5supja@localhost:5432/festivals-bdd?sslmode=disable" eu.gcr.io/kubernetessandbox-352207/festivals-api:0.0.1

.PHONY: postgres createdb dropdb image container