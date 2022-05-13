package data

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func GetFestivals(w http.ResponseWriter, r *http.Request) {

	jsonString :=
		`[
		 {
		   "id":"1",
		   "name":"festival 1",
		   "image":"https://user-images.githubusercontent.com/1500684/157774694-99820c51-8165-4908-a031-34fc371ac0d6.jpg",
		   "description":"Festival 1 description",
		   "isLiked":false,
		   "artist":[
			  "artist 1",
			  "artist 2"
		   ]
		},
		{
		   "id":"2",
		   "name":"festival 2",
		   "image":"https://user-images.githubusercontent.com/1500684/157774694-99820c51-8165-4908-a031-34fc371ac0d6.jpg",
		   "description":"Festival 2 description",
		   "isLiked":true,
		   "artist":[
			  "artist 3",
			  "artist 2"
		   ]
		},
		{
		   "id":"3",
		   "name":"festival 3",
		   "image":"https://user-images.githubusercontent.com/1500684/157774694-99820c51-8165-4908-a031-34fc371ac0d6.jpg",
		   "description":"Festival 2 description",
		   "isLiked":true,
		   "artist":[
			  "artist 3",
			  "artist 2"
		   ]
		},
		{
		   "id":"4",
		   "name":"festival 4",
		   "image":"https://user-images.githubusercontent.com/1500684/157774694-99820c51-8165-4908-a031-34fc371ac0d6.jpg",
		   "description":"Festival 2 description",
		   "isLiked":true,
		   "artist":[
			  "artist 3",
			  "artist 2"
		   ]
		},
		{
		   "id":"5",
		   "name":"festival 5",
		   "image":"https://user-images.githubusercontent.com/1500684/157774694-99820c51-8165-4908-a031-34fc371ac0d6.jpg",
		   "description":"Festival 2 description",
		   "isLiked":true,
		   "artist":[
			  "artist 3",
			  "artist 2"
		   ]
		},
		{
		   "id":"6",
		   "name":"festival 6",
		   "image":"https://user-images.githubusercontent.com/1500684/157774694-99820c51-8165-4908-a031-34fc371ac0d6.jpg",
		   "description":"Festival 6 description",
		   "isLiked":true,
		   "artist":[
			  "artist 3",
			  "artist 2"
		   ]
		},
		{
		   "id":"7",
		   "name":"festival 7",
		   "image":"https://user-images.githubusercontent.com/1500684/157774694-99820c51-8165-4908-a031-34fc371ac0d6.jpg",
		   "description":"Festival 7 description",
		   "isLiked":true,
		   "artist":[
			  "artist 3",
			  "artist 2"
		   ]
		},
		{
		   "id":"8",
		   "name":"festival 8",
		   "image":"https://user-images.githubusercontent.com/1500684/157774694-99820c51-8165-4908-a031-34fc371ac0d6.jpg",
		   "description":"Festival 2 description",
		   "isLiked":true,
		   "artist":[
			  "artist 3",
			  "artist 2"
		   ]
		},
		{
		   "id":"9",
		   "name":"festival 9",
		   "image":"https://user-images.githubusercontent.com/1500684/157774694-99820c51-8165-4908-a031-34fc371ac0d6.jpg",
		   "description":"Festival 2 description",
		   "isLiked":true,
		   "artist":[
			  "artist 3",
			  "artist 2"
		   ]
		},
		{
		   "id":"10",
		   "name":"festival 10",
		   "image":"https://user-images.githubusercontent.com/1500684/157774694-99820c51-8165-4908-a031-34fc371ac0d6.jpg",
		   "description":"Festival 10 description",
		   "isLiked":true,
		   "artist":[
			  "artist 3",
			  "artist 2"
		   ]
		}
	]`

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, []byte(jsonString), "", "\t")
	if error != nil {
		return
	}

	//data := Event{}
	//json.Unmarshal([]byte(jsonString), &data)

	w.WriteHeader(http.StatusCreated)
	//w.Write([]byte(prettyJSON))
}
