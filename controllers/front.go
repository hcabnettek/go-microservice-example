package controllers

import (
	"encoding/json"
	"io"
)

/*func RegisterControllers() {
	uc := NewUserController()

	http.HandleFunc("/users", uc.ServeHTTP)
	http.Handle("/users/", *uc)
}*/

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
