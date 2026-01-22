package handler

import "net/http"

type OrderHandlerDeps struct {
}

type OrderHandler struct {
}

func NewOrderHandler(router *http.ServeMux, deps OrderHandlerDeps) {
	handler := &OrderHandler{}

	// Routing
	router.HandleFunc("GET /order/{id}", handler.getOrderById())

}

func (handler *OrderHandler) getOrderById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
