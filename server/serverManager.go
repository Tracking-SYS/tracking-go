package server

import (
	"context"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"golang.org/x/sync/errgroup"

	"factory/exam/server/kafka"
	"factory/exam/utils/logger"
	"factory/exam/utils/shutdown"
	"factory/exam/utils/tracer"
)

//Manager ...
type Manager struct {
	httpServer    *HTTPServer
	metricServer  *MetricServer
	kafkaConsumer *kafka.KafkaConsumer
	kafkaProducer *kafka.KafkaProducer
	traceProvider *tracesdk.TracerProvider
}

//NewServerManager ...
func NewServerManager(
	httpServer *HTTPServer,
	metricServer *MetricServer,
	kafkaConsumer *kafka.KafkaConsumer,
	kafkaProducer *kafka.KafkaProducer,
) *Manager {
	return &Manager{
		httpServer:    httpServer,
		metricServer:  metricServer,
		kafkaConsumer: kafkaConsumer,
		kafkaProducer: kafkaProducer,
	}
}

//StartAll ...
func (m *Manager) StartAll(parentCtx context.Context) error {
	logger.InitLogger()
	m.traceProvider = tracer.InitTracer()
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

	//Start metric server on port 9992
	eg.Go(func() error {
		return shutdown.BlockListen(ctx, m.kafkaProducer)
	})

	return eg.Wait()
}

//CloseAll ...
func (m *Manager) CloseAll() error {
	sentry.Flush(2 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := m.traceProvider.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	return nil
}
