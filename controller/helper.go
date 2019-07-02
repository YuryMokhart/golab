package controller

import "net/http"

// TODO: all comments start with one space.
//ErrorHelper prints an error.
func ErrorHelper(w http.ResponseWriter, err error, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	// TODO: you don't process the w.Write error.
	w.Write([]byte(message + ` ` + err.Error()))
	return
}
