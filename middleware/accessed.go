package middleware

import (
	"net/http"

	"github.com/daichitakahashi/hrms"
)

// Accessed is a middleware
func Accessed(afterAccess func(*ResponseWriterWithStatusCode, *http.Request), recovered func(*ResponseWriterWithStatusCode, *http.Request, interface{})) func(hrms.Handler) hrms.Handler {
	return func(next hrms.Handler) hrms.Handler {
		fn := func(w http.ResponseWriter, r *http.Request, ps hrms.Params) {
			writer := NewResponseWriterWithStatusCode(w)

			defer func() {
				rvr := recover()
				if rvr != nil {
					recovered(writer, r, rvr)
				}
				afterAccess(writer, r)
			}()

			next.Handle(writer, r, ps)
		}
		return hrms.Handle(fn)
	}
}

// ResponseWriterWithStatusCode is wrapper
type ResponseWriterWithStatusCode struct {
	http.ResponseWriter
	StatusCode int
}

// NewResponseWriterWithStatusCode is
func NewResponseWriterWithStatusCode(w http.ResponseWriter) *ResponseWriterWithStatusCode {
	return &ResponseWriterWithStatusCode{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
	}
}

// WriteHeader is wrapper
func (r *ResponseWriterWithStatusCode) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
