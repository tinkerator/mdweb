# mdweb - a simple way to view markdown files locally.

## Overview

The `mdweb` program is a local webserver for serving up rendered `.md`
files and other content. It converts `.md` files to HTML in order to
render them with customized [fonts](/fonts.css) and [site](/site.css)
style (both of these files are embedded into the `mdweb` binary at
build time).

To use `mdweb` for the first time, just run these commands:
```
$ git clone https://github.com/tinkerator/mdweb.git
$ cd mdweb
$ go get
$ go build
$ ./mdweb
```
By default, `mdweb` listens to `localhost:8080` and serves files
rooted in the working directory. You can override these defaults with
the `--addr` and `--base` arguments respectively. As a first example,
when run from this present directory, point your browser to
http://localhost:8080/README.md.

## License info

The `mdweb` program is distributed with the same BSD 3-clause license
as that used by [golang](https://golang.org/LICENSE) itself.

## Reporting bugs and feature requests

The program `mdweb` has been developed purely out of self-interest and
a desire to quickly render `.md` files. If you find a bug or want to
suggest a feature addition, please use the [bug
tracker](https://github.com/tinkerator/mdweb/issues).
