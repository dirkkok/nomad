package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestSentinelReadCommand_Implements(t *testing.T) {
	t.Parallel()
	var _ cli.Command = &SentinelReadCommand{}
}
