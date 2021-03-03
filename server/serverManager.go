package server

import (
	"context"
	"time"

	"github.com/getsentry/sentry-go"
	"golang.org/x/sync/errgroup"

	"factory/exam/utils/logger"
	"factory/exam/utils/shutdown"
	"factory/exam/utils/tracer"
)

//Manager ...
type Manager struct {
	httpServer   *HTTPServer
	metricServer *MetricServer
	tracerFlush  func()
}

//NewServerManager ...
func NewServerManager(
	httpServer *HTTPServer,
	metricServer *MetricServer,
) *Manager {
	return &Manager{
		httpServer:   httpServer,
		metricServer: metricServer,
	}
}

//StartAll ...
func (m *Manager) StartAll(parentCtx context.Context) error {
	logger.InitLogger()
	m.tracerFlush = tracer.InitTracer()
	eg, ctx := errgroup.WithContext(parentCtx)

	//Start http server on port 8080
	eg.Go(func() error {
		return shutdown.BlockListen(ctx, m.httpServer)
	})

	//Start metric server on port 9992
	eg.Go(func() error {
		return shutdown.BlockListen(ctx, m.metricServer)
	})

	return eg.Wait()
}

//CloseAll ...
func (m *Manager) CloseAll() error {
	sentry.Flush(2 * time.Second)
	m.tracerFlush()
	return nil
}
