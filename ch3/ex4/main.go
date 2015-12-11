package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/chris-gilmore/gopl/ch3/ex4/surface"
)

func main() {
	var color string
	var width, height int

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			color = r.FormValue("color")
			if v, err := strconv.Atoi(r.FormValue("width")); err == nil {
				width = v
			}
			if v, err := strconv.Atoi(r.FormValue("height")); err == nil {
				height = v
			}

			w.Header().Set("Content-Type", "image/svg+xml")
			surface.SVG(w, color, width, height)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	surface.SVG(os.Stdout, color, width, height)
}
