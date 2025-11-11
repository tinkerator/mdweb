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
	_ "embed"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	mdh "github.com/gomarkdown/markdown/html"
)

var (
	addr     = flag.String("addr", "localhost:8080", "address to listen to for web request")
	base     = flag.String("base", "./", "base directory for html pages")
	renderer *mdh.Renderer
)

//go:embed css/site.css
var siteCSS []byte

//go:embed css/fonts.css
var fontsCSS []byte

var header = template.Must(template.New("name").Parse(`<!DOCTYPE html>
<html>
  <header>
    <meta charset="utf-8">
    <link rel="stylesheet" href="/fonts.css">
    <link rel="stylesheet" href="/site.css">
    <title>{{.}}</title>
  </header>
  <body>
`))

var footer = []byte(`  </body>
</html>`)

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if strings.Contains(path, "..") {
		http.NotFound(w, r)
		return
	}
	if strings.HasSuffix(path, ".md") {
		data, err := os.ReadFile(*base + path)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		header.Execute(w, path)
		w.Write(markdown.ToHTML(data, nil, renderer))
		w.Write(footer)
		return
	}
	switch path {
	case "/fonts.css":
		w.Header().Add("Content-Type", "text/css")
		w.Write(fontsCSS)
	case "/site.css":
		w.Header().Add("Content-Type", "text/css")
		w.Write(siteCSS)
	default:
		http.ServeFile(w, r, *base+path)
	}
}

func main() {
	flag.Parse()

	renderer = mdh.NewRenderer(mdh.RendererOptions{
		Flags: mdh.TOC,
	})
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
