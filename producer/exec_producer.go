package producer

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/go-task/task/execext"
)

// ExecProducer is the default producer for a task run
type ExecProducer struct {
}

// StartRun is called before the tasks are run
func (ep ExecProducer) StartRun() error {
	fmt.Println("Starting task")
	return nil
}

// RunTask is called for each task that is about to be executed
func (ep ExecProducer) RunTask(task string) error {
	fmt.Printf("\n*** %s ***\n\n", task)
	return nil
}

// RunCommand is called for each command of a task
func (ep ExecProducer) RunCommand(ctx context.Context, tSet string, opts execext.RunCommandOptions) error {
	fmt.Printf("--> %s\n", opts.Command)

	if tSet == "" {
		opts.Stdout = os.Stdout
		if err := execext.RunCommand(&opts); err != nil {
			return err
		}
	} else {
		buff := bytes.NewBuffer(nil)
		opts.Stdout = buff
		if err := execext.RunCommand(&opts); err != nil {
			return err
		}
		os.Setenv(tSet, strings.TrimSpace(buff.String()))
	}

	return nil
}

// FinishRun is called at the end
func (ep ExecProducer) FinishRun() error {
	fmt.Println("Finished task")
	return nil
}
