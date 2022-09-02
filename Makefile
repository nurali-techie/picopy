help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: clean					## build binary
	go build -o ./out/picopy .

install:						## install piccopy
	go install

clean:							## remove rc binary
	rm -f ./out/picopy

go-mod:							## run go modules tidy, verify
	go mod tidy
	go mod verify
