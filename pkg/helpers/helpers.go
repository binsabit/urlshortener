package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Enveleope map[string]interface{}

func ReadJSON(r *http.Request, dst interface{}) error {

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dst)
	if err != nil {
		return fmt.Errorf("Error json data decoding")
	}
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, data Enveleope, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}
