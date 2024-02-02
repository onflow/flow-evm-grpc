.PHONY: test
test:
	# test all packages
	go test -cover -parallel 8 ./...

.PHONY: check-tidy
check-tidy:
	go mod tidy
	git diff --exit-code

.PHONY: generate
generate:
	go get -d github.com/vektra/mockery/v2@v2.21.4
	mockery --all --dir=storage --output=storage/mocks
	mockery --all --dir=services/events --output=services/events/mocks

.PHONY: ci
ci: check-tidy test
