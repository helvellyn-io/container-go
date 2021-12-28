package main

import (
	"log"

	app "github.com/helvellyn-io/container-go"
	"github.com/helvellyn-io/container-go/src/git"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "builder", worker.Options{}) // builder task q needs to match the start q declaration

	w.RegisterWorkflow(app.BuildWorkflow)
	w.RegisterActivity(git.GetBuildArtifacts)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
