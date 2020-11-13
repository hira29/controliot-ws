package controller

import (
	"encoding/json"
	"net/http"
	"os"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var uname string
	uname = os.Getenv("DATABASE_UNAME")
	json.NewEncoder(w).Encode(uname)
}
