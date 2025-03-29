package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"regexp"
)

type JSONResponse struct {
	Error   string      `json:"error" bson:"error"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data,omitempty" bson:"data,omitempty"`
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, header ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(header) > 0 {
		for key, value := range header[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	// accept only 1mb of data
	maxbyte := 1024 * 1024
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxbyte))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(data)
	if err != nil {
		return err
	}

	// check for more than one request
	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("request body must only have a single JSON object")
	}

	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, message string, status ...int) error {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	var jr JSONResponse
	jr.Error = message
	jr.Message = message

	return app.writeJSON(w, statusCode, jr)
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
