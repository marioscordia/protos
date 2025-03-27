package http

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type jsonError struct {
	Message string `json:"error"`
}

func ResponseWithError(logger *slog.Logger, w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	je := jsonError{
		Message: msg,
	}
	if err := json.NewEncoder(w).Encode(je); err != nil {
		logger.Error("failed on encoding json error", slog.String("err", err.Error()))
	}
}

func Response(logger *slog.Logger, w http.ResponseWriter, msg interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		logger.Error("failed on encoding http msg", slog.String("err", err.Error()))
	}
}
