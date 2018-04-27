package middleware

import (
	"net/http"

	"github.com/daichitakahashi/hrms"
)

// accept characters for pattern matching
const (
	upper        = 'W'
	lower        = 'w'
	number       = 'n'
	hyphen       = '-'
	underscore   = '_'
	period       = '.'
	tilda        = '~'
	beginPattern = '('
	endPattern   = ')'
)

// ParamCheck is
func ParamCheck(patterns map[string]string) func(hrms.Handler) hrms.Handler {

	middleware := func(next hrms.Handler) hrms.Handler {
		fn := func(w http.ResponseWriter, r *http.Request, params hrms.Params) {

			for _, prm := range params {
				pattern, ok := patterns[prm.Key]
				if !ok {
					continue
				}
				//pattern check
				generateChecker(pattern)
			}

			next.Handle(w, r, params)
		}

		return hrms.Handle(fn)
	}

	return middleware
}

type checker struct {
	check int
}

func generateChecker(pattern string) bool {
	//
	return true
}
