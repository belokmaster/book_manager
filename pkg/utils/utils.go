package utils

//this package is parsing HTTP request and converting JSON data into a structure
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	// reads all data from r.body request body
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		// if everthing is good data is stored in the body variable.
		if err := json.Unmarshal([]byte(body), x); err != nil {
			// converts JSON data from the body to the structure pointed to x
			return
			// if an error occurs during reading or deserialization the function terminates execution
		}
	}
}
