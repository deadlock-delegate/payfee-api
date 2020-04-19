env:
	docker-compose -f etc/docker-compose.yml up -d db

migratestatus:
	go run cmd/migrate/main.go -dir db/migrations "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable" status

migratecreat:
	go run cmd/migrate/main.go -dir=db/migrations "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable" create $(NAME) sql

migrateup:
	go run cmd/migrate/main.go -dir=db/migrations "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable" up

migratedown:
	go run cmd/migrate/main.go -dir=db/migrations "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable" down

migratereset:
	go run cmd/migrate/main.go -dir=db/migrations "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable" reset
