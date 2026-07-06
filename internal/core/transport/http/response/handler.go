package core_http_response

import (
	"fmt"

	"encoding/json"
	"net/http"

	"go.uber.org/zap"
	core_logger "github.com/avequa/golang-todo-app/internal/core/logger"

)

type HTTPResponseHandler struct {
	log *core_logger.Logger
	rw http.ResponseWriter
}

func (h *HTTPResponseHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("panic: %v", p)

	h.log.Error(msg, zap.Error(err))
	h.rw.WriteHeader(statusCode)

	response := map[string]string{
		"message": msg,
		"error": err.Error(),
	}

	if err := json.NewEncoder(h.rw).Encode(response); err != nil {
		h.log.Error("http response", zap.Error(err))
	}
}

func NewHTTPResponseHandler(
	log *core_logger.Logger, 
	rw http.ResponseWriter,
) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log: log,
		rw: rw,
	}
}

