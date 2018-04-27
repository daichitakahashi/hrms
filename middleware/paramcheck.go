package middleware

import (
	"net/http"

	"github.com/daichitakahashi/httprouter"
)

// ParamCheck is
func ParamCheck() {}

// Example is
func Example(next httprouter.Handler) httprouter.Handler {
	fn := func(w http.ResponseWriter, r *http.Request, params httprouter.Params) { // これは Handle
		//
		next.Handle(w, r, params)
	}

	return httprouter.Handle(fn) //これはキャスト
}
