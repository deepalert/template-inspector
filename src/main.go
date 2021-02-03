package main

import (
	"context"
	"os"

	"github.com/deepalert/deepalert"
	"github.com/deepalert/deepalert/inspector"
	"github.com/m-mizutani/golambda"
)

var logger = golambda.Logger

type handler struct{}

func (x *handler) Callback(ctx context.Context, attr deepalert.Attribute) (*deepalert.TaskResult, error) {
	// FIXME: Write your inspection logic
	return nil, nil
}

func main() {
	golambda.Start(func(event golambda.Event) (interface{}, error) {
		bodies, err := event.DecapSNSonSQSMessage()
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

		hdlr := &handler{}

		logger.With("tasks", tasks).Info("Start inspection")
		if err := inspector.Start(inspector.Arguments{
			Context: event.Ctx,
			Tasks:   tasks,
			Handler: hdlr.Callback,
			// FIXME: Fill Author to your inspector name
			Author:          "???",
			FindingQueueURL: os.Getenv("FINDING_QUEUE_URL"),
			AttrQueueURL:    os.Getenv("ATTRIBUTE_QUEUE_URL"),
		}); err != nil {
			return nil, err
		}

		return nil, nil
	})
}
