all: get-artifacts build-with-artifacts push-artifacts test-build init-temporal-worker start-temporal-worker

help:      ## help for this makefile.    
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

get-artifacts: ##will attempt to download the build artifacts as defined in vars.json "artifactProvider".
	@go run ../cmd/main.go get-artifacts

build-with-artifacts: ##will attempt to build a container using the Dockerfile artifacts from --get-artifacts
	@go run ../cmd/main.go build-artifacts

push-artifacts:  ##will attempt to push the build container to the registry decalted in vars.josn "registry"
	to-do 

test-build:   ##runs all tests
	@go test ...

start-temporal-worker:
	@go run ../worker/main.go

start-temporal-task:
	@go run ../start/main.go

clean: ## deletes the downloded artifacts and built images. 
	@rm -rf .fromGit/* && rm -rf .fromGit
	@rm -rf ../worker/.fromGit
	@rm -rf ../cmd/.fromGit
