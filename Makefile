clean:
	go clean -i ./...

deps:
	glide install
	glide update

build_inside_docker:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vault-get .

build_osx:
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o vault-get .

build:
	docker-compose build

serve:
	docker-compose up
