main_package_path = .
binary_name = task-cli

.PHONE: test
test:
	go test ./...

.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...

.PHONY: build
build:
	go build -o ${binary_name} ${main_package_path}