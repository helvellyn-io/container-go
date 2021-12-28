package main

import (
	"flag"

	"github.com/helvellyn-io/container-go/src/docker"
	"github.com/helvellyn-io/container-go/src/git"
)

var operation *string

func init() {
	operation = flag.String("operation", "definition", "Gets build artifacts.")

}

func main() {

	flag.Parse()

	if *operation == "get-artifacts" {
		git.GetBuildArtifacts()
	} else if *operation == "build-container" {
		docker.BuildImage(docker.InitDockerBuild())
	}

}
