repo=$(shell basename "`pwd`")
gopher:
	@git init
	@go mod init github.com/mikejeuga/$(repo)
	@go mod tidy
	@touch .gitignore


t: test
test:
	@go test -v ./...


ut: unit-test
unit-test:
	@go test -v -tags=unit ./...

at: acceptance-test
acceptance-test:
	@docker-compose -f docker-compose.yml up -d
	@go test -v -tags=acceptance ./...
	@docker-compose down

run:
	@go run ./cmd/main.go

ic: init
init:
	@git add .
	@git commit -m "Initial commit"
	@git remote add origin git@github.com:mikejeuga/${repo}.git
	@git branch -M main
	@git push -u origin main

c: commit
commit:
	@git add .
	@git commit -am "$m"
	@git pull --rebase
	git push
