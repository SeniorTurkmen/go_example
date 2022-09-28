package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
)

func ParseBody(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return Unmarshal(body, v)
}

func Unmarshal(data []byte, p interface{}) error {
	err := json.Unmarshal(data, p)
	if err != nil {
		return err
	}

	fields := reflect.ValueOf(p).Elem()
	for i := 0; i < fields.NumField(); i++ {

		yourpojectTags := fields.Type().Field(i).Tag.Get("isRequired")
		if strings.Contains(yourpojectTags, "required") && fields.Field(i).IsZero() {
			return errors.New(fields.Type().Field(i).Name + " is missing")
		}

	}
	return nil
}

func ResponsWithError(w http.ResponseWriter, code int, msg string) {
	ResponsWithJson(w, code, map[string]string{"error": msg})
}

func ResponsWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		ResponsWithError(w, http.StatusInternalServerError, err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
