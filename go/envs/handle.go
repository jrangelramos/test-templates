package function
/*
This function template responds with a list of environment variables that starts with TEST_
The template is meant to be used in by `func config envs` e2e test
*/
import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "text/plain")
	testEnvVars := []string{}
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, "TEST_") {
			testEnvVars = append(testEnvVars, e)
		}
	}
	_, err := fmt.Fprintf(res, "%v\n", strings.Join(testEnvVars, "\n"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error or response write: %v", err)
	}
}
