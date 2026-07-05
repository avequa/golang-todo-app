package core_http_server

import "net/http"

type HTTPServer struct {
	mux *http.ServeMux
	config Config
	log *core_logger.Logger
}

func NewHTTPServer(
	config Config,
	log *core_logger.Logger,
) *HTTPServer {
	return &HTTPServer{
		mux: http.NewServeMux(),
		config: config,
		log: log,
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	server := &http.Server{
		Addr: h.config.Addr,
		Handler: h.mux,
	}

	ch := make(chan error, 1)

	go func() {

		defer close(ch)
		
		h.log.Warn("start http server", zap.String("addr", h.config.Addr))

		err := server.LestenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("server HTTP: %w", err)
		}
	case <- ctx.Done():
		h.log.Warn("shutdown http server")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			h.config.ShutdownTimeout,
		)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown http server: %w", err)
		}

		h.log.Warn("http server stop :)")
	}
	return nil
}