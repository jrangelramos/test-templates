package function

import (
	"context"
	"fmt"
	"function/resp"
	"net/http"
	"os"
)

func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "text/plain")
	res.WriteHeader(200)

	_, err := fmt.Fprintf(res, resp.Get("REMOTE_VALID")+"\n")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
	}
}
