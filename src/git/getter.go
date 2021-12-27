package git

//https://github.com/helvellyn-io/container-go/blob/main/build/Dockerfile

import (
	"context"
	"fmt"
	"os"

	getter "github.com/hashicorp/go-getter"
)

func GetBuildArtifacts() {

	client := &getter.Client{
		Ctx:  context.Background(),
		Dst:  ".fromGit",
		Dir:  true,
		Src:  "github.com/helvellyn-io/container-go/build/",
		Mode: getter.ClientModeDir,
		//define the type of detectors go getter should use, in this case only github is needed
		Detectors: []getter.Detector{
			&getter.GitHubDetector{},
		},
		//provide the getter needed to download the files
		Getters: map[string]getter.Getter{
			"git": &getter.GitGetter{},
		},
	}
	//download the files in build directory
	if err := client.Get(); err != nil {
		fmt.Fprintf(os.Stderr, "Error getting path %s: %v", client.Src, err)
		os.Exit(1)
	}

}
