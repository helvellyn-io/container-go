package main

import (
	"flag"
	"fmt"

	"github.com/helvellyn-io/container-go/src/docker"
	"github.com/helvellyn-io/container-go/src/git"
)

var getBuildArtifacts *string
var buildContainer *string
var flagSlice []string

func init() {
	getBuildArtifacts = flag.String("ga", "get-artifacts", "Gets build artifacts.")
	buildContainer = flag.String("bc", "build-container", "Builds a container using artifacts.")

}

func main() {

	flag.Parse()
	flagSlice = append(flagSlice, *getBuildArtifacts)
	flagSlice = append(flagSlice, *buildContainer)

loop:
	for _, v := range flagSlice {

		switch v {

		case "get-artifacts":
			//copy build artifacts from Git.
			git.GetBuildArtifacts()
			break loop

		case "build-container":
			//build container using build artifacts
			docker.BuildImage(docker.InitDockerBuild())
			break loop

		default:
			fmt.Println("Please supply an operation: get-artifacts , build-container")
			break loop
		}
	}
}
