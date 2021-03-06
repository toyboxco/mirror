package main

import (
	"context"
	"fmt"

	cmd "github.com/talismanco/mirror/pkg/cmd"
	config "github.com/talismanco/mirror/pkg/config"
	exit "github.com/talismanco/mirror/pkg/util/exit"
	osutil "github.com/talismanco/mirror/pkg/util/osutil"

	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) (err error) {
	var exists bool
	if exists, err = osutil.DirExists(config.Configuration.TemplateDirPath); !exists {
		if err = osutil.CreateDirs(config.Configuration.TemplateDirPath); err != nil {
			exit.Error(fmt.Errorf("Tried to initialize your template directory, but it has failed: %s", err))
		}
	} else if err != nil {
		exit.Error(fmt.Errorf("Failed to init: %s", err))
	}

	cmd.Run()

	return nil
}
