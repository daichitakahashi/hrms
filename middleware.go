package hrms

import "net/http"

// Handler is
type Handler interface {
	Handle(http.ResponseWriter, *http.Request, Params)
}

// Handle is
func (h Handle) Handle(w http.ResponseWriter, r *http.Request, params Params) {
	h(w, r, params)
}

// Use is
func (r *Router) Use(middlewares ...func(Handler) Handler) {
	r.middlewares = append(r.middlewares, middlewares...)
}

func (r *Router) applyMiddlewares(endpoint Handler) Handler {
	if len(r.middlewares) == 0 {
		return endpoint
	}
	i := len(r.middlewares) - 1
	folded := r.middlewares[i](endpoint)
	for i--; i >= 0; i-- {
		folded = r.middlewares[i](folded)
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
