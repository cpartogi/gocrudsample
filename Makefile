mock-gen:
	mockery --dir domain/tutorial --name TutorialUsecaseInterface --filename itutorial_usecase.go --output domain/tutorial/mocks --with-expecter

test:
	go test -v -cover ./...


coverage-report:
	go tool cover -html cover.out	

test-cov:
	go test -coverprofile cover.out ./...
	go tool cover -func cover.out	