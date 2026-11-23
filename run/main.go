package main

import (
	"os"

	"github.com/paketo-buildpacks/libreload-packit/watchexec"
	pnpmstart "github.com/goodrain/pnpm-start"
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func main() {
	logger := scribe.NewEmitter(os.Stdout).WithLevel(os.Getenv("BP_LOG_LEVEL"))

	reloader := watchexec.NewWatchexecReloader()

	packit.Run(
		pnpmstart.Detect(reloader),
		pnpmstart.Build(
			logger,
			reloader,
		),
	)
}
