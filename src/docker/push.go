package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func PushImage(client client.Client, image string) error {
	ctx := context.Background()

	client.ImagePush(ctx, "this_is_a_imagename:latest", types.ImagePushOptions{})

	return nil
}
