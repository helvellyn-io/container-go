package main

import (
	"github.com/helvellyn-io/container-go/src/docker"
	"github.com/helvellyn-io/container-go/src/git"
)

func main() {
	//artifact := flag.String("--feature", "--get-artifacts", "gets build artifacts from Git")
	//copy build artifacts from Git.
	git.GetBuildArtifacts()
	//build docker image using Dockerfile artifact.
	docker.BuildImage(docker.InitDockerBuild())

}
