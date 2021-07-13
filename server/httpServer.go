package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	httplib "github.com/Tracking-SYS/go-lib/http"
	services_pb "github.com/Tracking-SYS/proto-tracking-gen/go/tracking/services"

	"factory/exam/handler"
	"factory/exam/utils/gateway"
	"factory/exam/utils/shutdown"
)

var _ shutdown.ServerAbstract = &HTTPServer{}

type HTTPServer struct {
	server *http.Server
}

//HTTPProvider ...
func HTTPProvider(
	ctx context.Context,
	productHandler *handler.ProductHandler,
	prodPBHandler services_pb.ProductServiceServer,
) (*HTTPServer, error) {
	router := httplib.NewHTTPBuilder()

	gateway := runtime.NewServeMux(gateway.DefaultGateMuxOpts()...)
	err := services_pb.RegisterProductServiceHandlerServer(ctx, gateway, prodPBHandler)
	if err != nil {
		return nil, err
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome changes"))
	})

	router.Get("/products", productHandler.Get)
	router.Route("/", func(r chi.Router) {
		r.Mount("/v1", gateway)
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router.Build(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &HTTPServer{
		server: server,
	}, nil
}

//Start ...
func (s *HTTPServer) Start() error {
	return s.server.ListenAndServe()
}

//Close ...
func (s *HTTPServer) Close() error {
	return s.server.Shutdown(context.Background())
}
