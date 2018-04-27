package middleware

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/daichitakahashi/hrms"
)

// ParamRegexp provides parameter check
func ParamRegexp(patterns map[string]string, failedRedirect string) func(hrms.Handler) hrms.Handler {

	// for server: regexp.MustCompile here

	middleware := func(next hrms.Handler) hrms.Handler {
		fn := func(w http.ResponseWriter, r *http.Request, params hrms.Params) {

			for _, prm := range params {
				pattern, ok := patterns[prm.Key]
				if !ok {
					continue
				}

				// currently, this is for CGI...(for server program, we should make Regex instance when registered this middleware)
				regex, err := regexp.Compile(pattern)
				if err != nil {
					fmt.Println(w, "regex error")
				}
				if !regex.MatchString(prm.Value) {
					http.Redirect(w, r, failedRedirect, 301) //http.StatusSeeOther)
				}

			}

			next.Handle(w, r, params)
		}

		return hrms.Handle(fn)
	}

	return middleware
}
