all: get-artifacts build-with-artifacts push-artifacts test-build init-temporal-worker start-temporal-worker

help:      ## help for this makefile.    
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

get-artifacts: ##will attempt to download the build artifacts as defined in vars.json "artifactProvider".
	@go run ../cmd/main.go
build-with-artifacts: ##will attempt to build a container using the Dockerfile artifacts from --get-artifacts
	touch two
push-artifacts:  ##will attempt to push the build container to the registry decalted in vars.josn "registry"
	touch three
test-build:   ##runs all tests
	@go test ...

clean: ## deletes the downloded artifacts and built images. 
	@rm -f .fromGit/*
