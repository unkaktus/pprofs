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
	_ "net/http/pprof"
)

func init() {
	// Save defaultServeMux already used by pprof
	pprofServeMux := http.DefaultServeMux
	// Overwrite defaultServeMux with a clean one
	http.DefaultServeMux = http.NewServeMux()
	go func() {
		err := http.ListenAndServe(":6060", pprofServeMux)
		if err != nil {
			log.Printf("error serving pprof: %v", err)
		}
	}()
}
