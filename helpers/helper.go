package helpers

import (
	"fmt"
	"net/http"
)

// TODO: that function is a part of http layer. Your http layer is small, so you can just keep that function in http layer.
// ErrorHelper prints an error.
func ErrorHelper(w http.ResponseWriter, err error, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	// TODO: think about the better solution.
	_, err = w.Write([]byte(message + `: ` + err.Error()))
	if err != nil {
		fmt.Printf("could not write a header in wire format %v", err)
		return
	}
	return
}
