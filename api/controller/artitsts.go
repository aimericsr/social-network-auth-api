package controller

import (
	"encoding/json"
	"net/http"

	"github.com/aimericsr/social-network-auth-api/db"
)

func GetArtists(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetArtists()

	if err != nil {
		w.Write([]byte("kdf"))

	}

	_, err = json.Marshal(users)
	if err != nil {

	}

	//w.Write([]byte(res))
}
