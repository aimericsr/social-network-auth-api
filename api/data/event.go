package data

type Event struct {
	ID          int      `json:id`
	Name        string   `json:name`
	Image       string   `json:image`
	Description string   `json:description`
	IsLiked     bool     `json:isLiked`
	Artist      []string `json:artist`
}
