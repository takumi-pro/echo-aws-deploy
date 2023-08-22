.PHONY: go-up
go-up:
	docker-compose -f docker-compose-golang.yml up -d

.PHONY: go-down
go-down:
	docker-compose -f docker-compose-golang.yml down

.PHONY: go-rebuild-up
go-rebuild-up:
	docker-compose -f docker-compose-golang.yml up -d --build