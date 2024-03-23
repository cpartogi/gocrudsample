mock-gen:
	mockery --dir domain/tutorial --name TutorialUsecaseInterface --filename itutorial_usecase.go --output domain/tutorial/mocks --with-expecter
	mockery --dir domain/tutorial --name TutorialRepoInterface --filename itutorial_repo.go --output domain/tutorial/mocks --with-expecter

coverage-report:
	go tool cover -html cover.out	

test:
	go test -coverprofile cover.out ./...
	go tool cover -func cover.out	

run:	
	go run main.go

docker-restart:
	docker compose down -v
	docker compose up -d

migration-up:
	migrate -path migration/sql -database "postgresql://postgres:postgres@localhost:6432/postgres?sslmode=disable" -verbose up