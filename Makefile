docker-up:
	docker-compose up -d
psql:
	docker-compose exec db psql -U template
test_psql:
	docker-compose exec test_db psql -U test
run:
	./script/run.sh
test:
	./script/test.sh
.PHONY: mock
mock:
	script/mock.sh
schema:
	go generate ./ent
.PHONY: migration
migration:
	script/migration.sh
.PHONY: proto
proto:
	protoc ./proto/template/**/*.proto --go_out=./proto/ --go-grpc_out=./proto/
