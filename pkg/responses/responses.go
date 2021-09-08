package responses

import (
	"encoding/json"
	"net/http"
)

// WriteData writes the data onto the response writer.
func WriteData(w http.ResponseWriter, data interface{}) {
	var res Response
	res.Data = data
	json.NewEncoder(w).Encode(res)
}

func WriteError(w http.ResponseWriter, err error) {
	var res Response
	res.Error = err.Error()
	json.NewEncoder(w).Encode(res)
}
