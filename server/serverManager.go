package server

import (
	"context"
	"time"

	"github.com/getsentry/sentry-go"
	"golang.org/x/sync/errgroup"

	"github.com/Tracking-SYS/tracking-go/utils/logger"
	"github.com/Tracking-SYS/tracking-go/utils/shutdown"
	"github.com/Tracking-SYS/tracking-go/utils/tracer"
)

//Manager ...
type Manager struct {
	httpServer    *HTTPServer
	metricServer  *MetricServer
	kafkaConsumer *KafkaConsumer
	tracerFlush   func()
}

//NewServerManager ...
func NewServerManager(
	httpServer *HTTPServer,
	metricServer *MetricServer,
	kafkaConsumer *KafkaConsumer,
) *Manager {
	return &Manager{
		httpServer:    httpServer,
		metricServer:  metricServer,
		kafkaConsumer: kafkaConsumer,
	}
}

//StartAll ...
func (m *Manager) StartAll(parentCtx context.Context) error {
	logger.InitLogger()
	m.tracerFlush = tracer.InitTracer("tracking.sys.go")
	eg, ctx := errgroup.WithContext(parentCtx)

	//Start http server on port 8080
	eg.Go(func() error {
		return shutdown.BlockListen(ctx, m.httpServer)
	})

	//Start metric server on port 9992
	eg.Go(func() error {
		return shutdown.BlockListen(ctx, m.metricServer)
	})

	//Start metric server on port 9992
	eg.Go(func() error {
		return shutdown.BlockListen(ctx, m.kafkaConsumer)
	})

	return eg.Wait()
}

//CloseAll ...
func (m *Manager) CloseAll() error {
	sentry.Flush(2 * time.Second)
	m.tracerFlush()
	return nil
}
