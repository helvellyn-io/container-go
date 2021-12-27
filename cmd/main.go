package main

import (
	"log"

	"github.com/docker/docker/client"
	"github.com/helvellyn-io/container-go/src/docker"
)

func main() {
	client, err := client.NewClientWithOpts()
	if err != nil {
		log.Fatalf("Unable to create docker client: %s", err)
	}

	tags := []string{"this_is_a_imagename"}
	dockerfile := "/Users/dylanjohnson/go/src/github.com/helvellyn-io/container-go/build/Dockerfile"
	err = docker.BuildImage(*client, tags, dockerfile)
	if err != nil {
		log.Println(err)
	}
}
