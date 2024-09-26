package function

import (
	"fmt"
	"function/resp"
	"net/http"
	"os"
)

// Handle an HTTP Request.
func Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(200)
	_, err := fmt.Fprintf(w, resp.Get("REMOTE_VALID")+"\n")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
	}
}
