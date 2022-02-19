package worker

import (
	"context"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
)

// Worker ...
type Worker struct {
	handlers map[string]HandlerFunc
}

// Interface ...
type Interface interface {
	AddHandler(name string, handler HandlerFunc)
	Start()
}

// NewWorker initializes worker
func NewWorker() Interface {
	return &Worker{
		handlers: make(map[string]HandlerFunc),
	}
}

// AddHandler appends handler to the worker instance
func (w *Worker) AddHandler(name string, handler HandlerFunc) {
	w.handlers[name] = handler
}

// Start starts worker
func (w *Worker) Start() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	for _, handler := range w.handlers {
		go func(ctx context.Context, handler HandlerFunc) {
			handler.Start(ctx)
		}(ctx, handler)
	}

	waitForCancelSignal(ctx, cancel)
}

// waitForCancelSignal will check for cancel signal from the context
func waitForCancelSignal(ctx context.Context, cancelFunc context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for range c {
		logrus.WithField("signal", "interrupt").Info("[worker] received interrupt signal")
		cancelFunc()
		break
	}
}
