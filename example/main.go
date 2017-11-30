// main.go - example program to use pprofs
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to this module of pprofs, using the creative
// commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package main

import (
	"log"
	"net/http"

	_ "github.com/nogoegst/pprofs"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hey"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
