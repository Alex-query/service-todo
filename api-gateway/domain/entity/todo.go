package entity

type Todo struct {
	Id      string  `json:"id"`
	Title   string `json:"title"`
	OwnerId int64  `json:"owner_id"`
}
