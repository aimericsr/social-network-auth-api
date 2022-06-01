package socialnetwork

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aimericsr/social-network-auth-api/api"
	"github.com/aimericsr/social-network-auth-api/model"
	"github.com/aimericsr/social-network-auth-api/util"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var user model.User
	resJson, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Can't Unmarshal the body", err)
	}
	err = json.Unmarshal([]byte(resJson), &user)

	hashedPassword, err := util.HashPassword(user.Password)

	dbRes := api.Serv.DB.Create(&model.User{Username: user.Username, Password: hashedPassword, Email: user.Email, Description: user.Description})
	if dbRes.Error != nil {
		fmt.Printf("err in database %v\n", dbRes.Error)
		//log.Fatal("pb", dbRes.Error)
	}
	w.WriteHeader(http.StatusCreated)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req loginUserRequest
	resJson, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Can't Unmarshal the body", err)
	}
	err = json.Unmarshal([]byte(resJson), &req)

	user := model.User{Username: req.Username}
	dbRes := api.Serv.DB.First(&user)
	if dbRes.Error != nil {
		fmt.Printf("err in database %v\n", dbRes.Error)
		//log.Fatal("pb", dbRes.Error)
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		fmt.Printf("err in database %v\n", dbRes.Error)
	}

	// accessToken, accessPayload, err := server.Serv.TokenMaker.CreateToken(
	// 	user.Username,
	// 	server.Serv.Config.AccessTokenDuration,
	// )
	// if err != nil {
	// 	fmt.Printf("err in database %v\n", dbRes.Error)
	// }

}
