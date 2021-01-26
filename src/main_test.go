package main_test

import (
	"context"
	"testing"

	"github.com/deepalert/deepalert"
	"github.com/deepalert/deepalert/inspector"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	main "github.com/deepalert/template-inspector/src"
)

func TestInspectorExample(t *testing.T) {
	attrURL := "https://sqs.ap-northeast-1.amazonaws.com/123456789xxx/attribute-queue"
	findingURL := "https://sqs.ap-northeast-1.amazonaws.com/123456789xxx/content-queue"

	hdlr := &main.Handler{}

	args := inspector.Arguments{
		Context:         context.Background(),
		Handler:         hdlr.Callback,
		Author:          "blue",
		AttrQueueURL:    attrURL,
		FindingQueueURL: findingURL,
	}

	t.Run("With attribute", func(tt *testing.T) {
		mock, newSQS := inspector.NewSQSMock()
		args.NewSQS = newSQS

		task := &deepalert.Task{
			ReportID: deepalert.ReportID(uuid.New().String()),
			// TODO: Add attribute to be inspected
			Attribute: &deepalert.Attribute{
				// Example:
				Type:  deepalert.TypeIPAddr,
				Key:   "dst",
				Value: "192.10.0.1",
			},
		}

		err := inspector.HandleTask(context.Background(), task, args)
		require.NoError(tt, err)
		sections, err := mock.GetSections(findingURL)
		require.NoError(tt, err)

		// len(setions) == 1 means 1 finding is returned
		require.Equal(tt, 0, len(sections))
	})
}
