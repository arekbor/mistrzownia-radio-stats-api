package utils

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"
)

func WriteJSON(w http.ResponseWriter, v any) error {
	res, err := json.Marshal(v)
	if err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(res)
	return err
}

func GetMaxDbTimeout() (time.Duration, error) {
	sec, err := strconv.Atoi(os.Getenv("DB_MAX_TIMEOUT"))
	if err != nil {
		return 0, err
	}
	dur := time.Duration(sec) * time.Second
	return dur, nil
}
