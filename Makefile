PWD = $(shell pwd)

# Запустить тесты
.PHONY: test
test:
	go test $(PWD)/... -coverprofile=cover.out

