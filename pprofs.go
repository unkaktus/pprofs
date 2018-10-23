// pprofs.go - serve pprof in a safe way
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to this module of pprofs, using the creative
// commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package pprofs

import (
	"log"
	"net/http"

	"github.com/nogoegst/pprofs/default-mux-saver"

	_ "net/http/pprof"
)

func init() {
	// Save DefaultServeMux already populated by pprof
	pprofServeMux := http.DefaultServeMux
	// Restore original DefaultServeMux
	default_mux_saver.Restore()
	go func() {
		err := http.ListenAndServe(":6060", pprofServeMux)
		if err != nil {
			log.Printf("error serving pprof: %v", err)
		}
	}()
}
