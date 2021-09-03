package server

import (
	"context"
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/go-chi/chi"
	otel_global "go.opentelemetry.io/otel/metric/global"

	prometheusLib "github.com/Tracking-SYS/go-lib/prometheus"

	"github.com/Tracking-SYS/tracking-go/config"
)

//MetricServer ...
type MetricServer struct {
	server *http.Server
}

//NewMetricServer ...
func NewMetricServer(
	metricPort config.MetricPort,
) (*MetricServer, error) {
	exporter, err := prometheusLib.InitExporter()
	if err != nil {
		return nil, fmt.Errorf("init exporter: %w", err)
	}

	if err := prometheusLib.WithRuntime(); err != nil {
		return nil, fmt.Errorf("init runtimemetrics: %w", err)
	}

	otel_global.SetMeterProvider(exporter.MeterProvider())

	mux := chi.NewRouter()
	mux.Get("/metrics", exporter.ServeHTTP)

	if os.Getenv("ENABLE_PPROF") == "yes" {
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
		mux.HandleFunc("/debug/pprof/*", pprof.Index)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", metricPort),
		Handler: mux,
	}
	return &MetricServer{server: server}, nil
}

//Start ...
func (s *MetricServer) Start() error {
	return s.server.ListenAndServe()
}

//Close ...
func (s *MetricServer) Close() error {
	err := s.server.Shutdown(context.Background())
	if err != nil {
		fmt.Println("Close MetricServer has error: ", err)
		return err
	}

	return nil
}
