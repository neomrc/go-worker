package worker

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	cb       func()
	interval int
}

type HandlerFunc interface {
	Start(ctx context.Context)
	SetInterval(interval int)
}

func NewHandler(cb func()) HandlerFunc {
	return &Handler{
		cb: cb,
	}
}

func (h *Handler) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Duration(h.interval) * time.Millisecond):
			go func() {
				t := time.Now()
				defer logrus.WithField("duration", time.Since(t)).WithField("interval", h.interval).Info("[worker] execution time for task")
				h.cb()
			}()
		}
	}
}

func (h *Handler) SetInterval(interval int) {
	h.interval = interval
}
