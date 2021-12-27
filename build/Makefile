all: --get-artifacts --build-with-artifacts --push-artifacts --test-build --init-temporal-worker --start-temporal-worker

help:      ## help for this makefile.    
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

--get-artifacts: ##will attempt to download the build artifacts as defined in vars.json "artifactProvider".
	downloading build artifacts .. 
	@go run ../cmd/main.go
	done .. 
--build-with-artifacts: ##will attempt to build a container using the Dockerfile artifacts from --get-artifacts
	touch two
--push-artifacts:  ##will attempt to push the build container to the registry decalted in vars.josn "registry"
	touch three
--test-build:   ##runs all tests
	running tests
	@go test ...
	done ..
--init-temporal-worker: ## submits a temporal worker
	@todo

--start-temporal-worker: ## starts the temporal worker 
	@todo

--clean: ## deletes the downloded artifacts and built images. 
	deleting build artifacts .. 
	@rm -f .fromGit/*
	done .. 