package main

import (
	"context"
	"os"

	"github.com/deepalert/deepalert"
	"github.com/deepalert/deepalert/inspector"
	"github.com/m-mizutani/golambda"
)

var logger = golambda.Logger

// Handler is exported for main_test
func Handler(ctx context.Context, attr deepalert.Attribute) (*deepalert.TaskResult, error) {
	// TODO: Write your inspection logic
	return &deepalert.TaskResult{}, nil
}

func main() {
	golambda.Start(func(event golambda.Event) (interface{}, error) {
		bodies, err := event.DecapSQSBody()
		if err != nil {
			return nil, err
		}

		var tasks []*deepalert.Task
		for _, body := range bodies {
			var task deepalert.Task
			if err := body.Bind(&task); err != nil {
				return nil, err
			}
			tasks = append(tasks, &task)
		}

		logger.With("tasks", tasks).Info("Start inspection")
		if err := inspector.Start(inspector.Arguments{
			Context: event.Ctx,
			Tasks:   tasks,
			Handler: Handler,
			// TODO: Fill Author to your inspector name
			Author:          "???",
			FindingQueueURL: os.Getenv("FINDING_QUEUE_URL"),
			AttrQueueURL:    os.Getenv("ATTRIBUTE_QUEUE_URL"),
		}); err != nil {
			return nil, err
		}

		return nil, nil
	})
}
