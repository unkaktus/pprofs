// saver.go - http.DefaultServeMux saver
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to this module of pprofs, using the creative
// commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package default_mux_saver

import "net/http"

var savedMux *http.ServeMux

func init() {
	savedMux = http.DefaultServeMux
	http.DefaultServeMux = http.NewServeMux()
}

func Restore() {
	http.DefaultServeMux = savedMux
}
