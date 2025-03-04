package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/DmitriiSvarovskii/booking-service-go/internal/app/config"
)

type Handler struct {
	cfg *config.AppConfig
}

func NewHandler(cfg *config.AppConfig) *Handler {
	return &Handler{cfg: cfg}
}

func (h *Handler) BaseURL(rw http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(body))
}

func (h *Handler) MethodNotAllowedHandle(rw http.ResponseWriter, r *http.Request) {
	responseMessage := fmt.Sprintf("The method '%s' is not allowed for path '%s'.", r.Method, r.URL.Path)
	rw.WriteHeader(http.StatusMethodNotAllowed)
	io.WriteString(rw, responseMessage)
}
