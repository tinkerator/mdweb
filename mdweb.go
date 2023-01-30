// Program mdweb is an HTML web server renderer with special handling
// for .md file content (it renders it in HTML).
//
// To use it, just run this command with the working directory equal
// to the file tree you want it to serve data from.
//
//	$ cd html
//	$ mdweb
//
// By default, mdweb listens to localhost:8080. You can override this
// default with the --addr argument.
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
)

var addr = flag.String("addr", "localhost:8080", "address to listen to for web request")

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if strings.Contains(path, "..") {
		http.NotFound(w, r)
		return
	} else if strings.HasSuffix(path, ".md") {
		data, err := os.ReadFile("./" + path)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		w.Write(markdown.ToHTML(data, nil, nil))
	} else {
		http.ServeFile(w, r, "./"+path)
	}
}

func main() {
	flag.Parse()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
