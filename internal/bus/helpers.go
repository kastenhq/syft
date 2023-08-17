package bus

import (
	"github.com/wagoodman/go-partybus"

	"github.com/kastenhq/syft/internal/log"
	"github.com/kastenhq/syft/syft/event"
)

func Exit() {
	Publish(partybus.Event{
		Type: event.CLIExit,
	})
}

func Report(report string) {
	if len(report) == 0 {
		return
	}
	report = log.Redactor.RedactString(report)
	Publish(partybus.Event{
		Type:  event.CLIReport,
		Value: report,
	})
}

func Notify(message string) {
	Publish(partybus.Event{
		Type:  event.CLINotification,
		Value: message,
	})
}
