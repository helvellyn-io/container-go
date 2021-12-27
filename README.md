# container-go
[![Go](https://github.com/helvellyn-io/container-go/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/helvellyn-io/container-go/actions/workflows/build.yml)


Prototyping continuous integration for microservices using Golang with Temporal as the build orchestration platform.

- For use with Temporal build task queues. 
  For guidance on how to run this as a Temporal task please refer to: 
  * https://github.com/helvellyn-io/temporal-weather

*MAKEFILE 
```
help:       help for this makefile.    
--get-artifacts: will attempt to download the build artifacts as defined in vars.json "artifactProvider".
--build-with-artifacts: will attempt to build a container using the Dockerfile artifacts from --get-artifacts
--push-artifacts:  will attempt to push the build container to the registry decalted in vars.josn "registry"
--test-build:   runs all tests
--init-temporal-worker:  submits a temporal worker
--start-temporal-worker:  starts the temporal worker 
--clean:  deletes the downloded artifacts, assocated images & terminates all temporal tasks.
```


