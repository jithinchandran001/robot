.DEFAULT_GOAL := greeting
.PHONY: greeting start test composer-up composer-down migration

greeting:
	@echo "make greeting start test composer-up"

test:
	go clean -testcache && export APP_ENV=dev && go mod vendor && cd tests && go test -mod=vendor -v . && cd ..

composer-up:
	- docker-compose kill
	docker-compose up -d
	@echo "wait for postgres container to be ready to establish connection"

composer-down:
	docker-compose kill
	#docker-compose rm -f

start:
	go run main.go

migration:
ifneq ($(and $(DBUSER),$(DBPASS),$(DBNAME),$(DBPORT)),)
	goose -dir migration postgres "user=$(DBUSER) password=$(DBPASS) dbname=$(DBNAME) sslmode=disable port=$(DBPORT)" up
else
	goose -dir migration postgres "user=postgres password=postgres dbname=robot sslmode=disable port=5432" up
endif
