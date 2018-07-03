package hrms

import "net/http"

// Handler is
type Handler interface {
	Handle(http.ResponseWriter, *http.Request, Params)
}

// HttpHandler is
func HttpHandler(h func(http.ResponseWriter, *http.Request)) Handle {
	return func(w http.ResponseWriter, r *http.Request, _ Params) {
		h(w, r)
	}
}

// Handle is
func (h Handle) Handle(w http.ResponseWriter, r *http.Request, params Params) {
	h(w, r, params)
}

// Use is
func (r *Router) Use(middlewares ...func(Handler) Handler) {
	r.middlewares = append(r.middlewares, middlewares...)
}

func applyMiddlewares(endpoint Handler, middlewares []func(Handler) Handler) Handler {
	if middlewares == nil || len(middlewares) == 0 {
		return endpoint
	}
	i := len(middlewares) - 1
	folded := middlewares[i](endpoint)
	for i--; i >= 0; i-- {
		folded = middlewares[i](folded)
	}

	return folded
}

// Example is
func Example(next Handler) Handler {
	fn := func(w http.ResponseWriter, r *http.Request, params Params) { // これは Handle
		//
		next.Handle(w, r, params)
	}

	return Handle(fn) //これはキャスト
}
