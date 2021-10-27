package function
/*
This function template returns on body the content of a file on the server
The template is meant to be used in by `func config volumes` e2e test
*/
import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"net/http"
	"strings"
)

func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "text/plain")

	// v=/test/volume-config/myconfig
	q := strings.Split(req.URL.RawQuery, "=")
	action := q[0]
	path := q[1]

	if action == "v" {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading file: %v", err)
		}
		_, err = fmt.Fprintf(res, "%v", content)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error on response write: %v", err)
		}
	}

}
