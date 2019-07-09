package helpers

import "net/http"

// ErrorHelper prints an error.
func ErrorHelper(w http.ResponseWriter, err error, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	_, error := w.Write([]byte(message + ` ` + err.Error()))
	if error != nil {
		// find how to process error in this case.
		return
	}
	return
}
