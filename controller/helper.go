package controller

import "net/http"

//ErrorHelper prints an error.
func ErrorHelper(w http.ResponseWriter, err error, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(message + ` ` + err.Error()))
	return
}
