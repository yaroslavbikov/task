package task

import (
	"context"

	"github.com/go-task/task/execext"
	"github.com/go-task/task/producer"
)

var (
	currentProducer Producer
)

// Producer executes a task definition
type Producer interface {
	StartRun() error
	RunTask(task string) error
	RunCommand(ctx context.Context, tSet string, opts execext.RunCommandOptions) error
	FinishRun() error
}

func init() {
	currentProducer = producer.ExecProducer{}
}
