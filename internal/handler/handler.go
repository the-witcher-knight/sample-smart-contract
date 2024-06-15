package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/the-witcher-knight/store-contract/internal/controller"
)

type Handler struct {
	ctrl controller.Controller
}

func New(ctrl controller.Controller) Handler {
	return Handler{ctrl: ctrl}
}

func (h Handler) Store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		valStr := chi.URLParam(r, "val")
		val, err := strconv.Atoi(valStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if err := h.ctrl.Store(r.Context(), int64(val)); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h Handler) Retrieve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		val, err := h.ctrl.Retrieve(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, val)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
