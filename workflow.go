package worker

import (
	"time"

	"github.com/helvellyn-io/container-go/src/git"
	"go.temporal.io/sdk/workflow"
)

// Workflow is a Builder  workflow definition.
func BuildWorkflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Git get artifacts workflow started", "name", name)

	var result string
	err := workflow.ExecuteActivity(ctx, git.GetBuildArtifacts, name).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	logger.Info("Git get artifacts workflow completed.", "result", result)

	return result, nil
}
