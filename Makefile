all:
	gofmt -e -s -w ${GOPATH}/src/github.com/sankalpjonn/wakanda
	go vet -v github.com/sankalpjonn/wakanda
	go get -v github.com/sankalpjonn/wakanda
	go install github.com/sankalpjonn/wakanda
	exit 0
