package api

import (
	"fmt"
	"net/http"

	"github.com/binsabit/urlshortener/pkg/helpers"
)

func (app *application) CreateURL(w http.ResponseWriter, r *http.Request) {
	var input struct {
		URL string `json:"url"`
	}

	err := helpers.ReadJSON(r, &input)

	if err != nil {
		http.NotFound(w, r)
	}

	//create it

	//test for error

	//return response success/fail
}

func (app *application) GoToURL(w http.ResponseWriter, r *http.Request) {
	//match short url with long
	//forward to long
}

func (app *application) DeleteURL(w http.ResponseWriter, r *http.Request) {
	//parse url
	//delete url
	//test for error
	//send response succcess/fail
}

func (app *application) Main(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
