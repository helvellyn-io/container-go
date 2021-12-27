package main

import (
	"log"

	"github.com/docker/docker/client"
	"github.com/helvellyn-io/container-go/src/docker"
	"github.com/helvellyn-io/container-go/src/git"
)

func main() {
	//copy build artifacts from Git.
	git.GetBuildArtifacts()

	//create docker client
	client, err := client.NewClientWithOpts()
	if err != nil {
		log.Fatalf("Unable to create docker client: %s", err)
	}
	//tags etc.
	tags := []string{"this_is_a_imagename"}
	dockerfile := "/Users/dylanjohnson/go/src/github.com/helvellyn-io/container-go/build/Dockerfile"
	//build
	err = docker.BuildImage(*client, tags, dockerfile)
	if err != nil {
		log.Println(err)
	}

}
