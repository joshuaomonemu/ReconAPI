package handlers

import "net/http"

func Appenduser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("firstname", "voke")
	w.Header().Set("lastname", "omonemu")

}
