package responses

import (
	"encoding/json"
	"net/http"
)

// WriteData writes the data onto the response writer, the goal is to format the responses. 
func WriteData(w http.ResponseWriter, data interface{}) {
	var res Response
	res.Data = data
	json.NewEncoder(w).Encode(res)
}

// WriteData writes the error onto the response writer, the goal is to format the responses. 
func WriteError(w http.ResponseWriter, err error) {
	var res Response
	res.Error = err.Error()
	json.NewEncoder(w).Encode(res)
}
