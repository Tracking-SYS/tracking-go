package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	errorsLib "github.com/Tracking-SYS/go-lib/errors"
	httplib "github.com/Tracking-SYS/go-lib/http"
	services_pb "github.com/Tracking-SYS/proto-tracking-gen/go/tracking/services"

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
	prodPBHandler services_pb.ProductServiceServer,
	taskPBHandler services_pb.TaskServiceServer,
) (*HTTPServer, error) {
	router := httplib.NewHTTPBuilder()

	gateway := runtime.NewServeMux(gateway.DefaultGateMuxOpts()...)
	err := errorsLib.ErrAny(
		services_pb.RegisterProductServiceHandlerServer(ctx, gateway, prodPBHandler),
		services_pb.RegisterTaskServiceHandlerServer(ctx, gateway, taskPBHandler),
	)

	if err != nil {
		return nil, err
	}

	userMetaMiddleware := httplib.NewRequestExtractor(
		httplib.NewRequestExtractorConf(httplib.CtxUserMetadata, httplib.ExtractUserMetaFromRequest),
	)

	router.Route("/", func(r chi.Router) {
		r.Use(userMetaMiddleware)
		r.Mount("/v1", gateway)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome changes"))
		})
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
