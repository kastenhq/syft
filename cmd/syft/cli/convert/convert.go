package convert

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/wagoodman/go-partybus"

	"github.com/anchore/stereoscope"
	"github.com/kastenhq/syft/cmd/syft/cli/eventloop"
	"github.com/kastenhq/syft/cmd/syft/cli/options"
	"github.com/kastenhq/syft/cmd/syft/internal/ui"
	"github.com/kastenhq/syft/internal/bus"
	"github.com/kastenhq/syft/internal/config"
	"github.com/kastenhq/syft/internal/log"
	"github.com/kastenhq/syft/syft"
	"github.com/kastenhq/syft/syft/formats"
	"github.com/kastenhq/syft/syft/sbom"
)

func Run(_ context.Context, app *config.Application, args []string) error {
	log.Warn("convert is an experimental feature, run `syft convert -h` for help")

	writer, err := options.MakeSBOMWriter(app.Outputs, app.File, app.OutputTemplatePath)
	if err != nil {
		return err
	}

	// could be an image or a directory, with or without a scheme
	userInput := args[0]

	var reader io.ReadCloser

	if userInput == "-" {
		reader = os.Stdin
	} else {
		f, err := os.Open(userInput)
		if err != nil {
			return fmt.Errorf("failed to open SBOM file: %w", err)
		}
		defer func() {
			_ = f.Close()
		}()
		reader = f
	}

	eventBus := partybus.NewBus()
	stereoscope.SetBus(eventBus)
	syft.SetBus(eventBus)
	subscription := eventBus.Subscribe()

	return eventloop.EventLoop(
		execWorker(reader, writer),
		eventloop.SetupSignals(),
		subscription,
		stereoscope.Cleanup,
		ui.Select(options.IsVerbose(app), app.Quiet)...,
	)
}

func execWorker(reader io.Reader, writer sbom.Writer) <-chan error {
	errs := make(chan error)
	go func() {
		defer close(errs)
		defer bus.Exit()

		s, _, err := formats.Decode(reader)
		if err != nil {
			errs <- fmt.Errorf("failed to decode SBOM: %w", err)
			return
		}

		if s == nil {
			errs <- fmt.Errorf("no SBOM produced")
			return
		}

		if err := writer.Write(*s); err != nil {
			errs <- fmt.Errorf("failed to write SBOM: %w", err)
		}
	}()
	return errs
}
