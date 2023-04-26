package http

type Attach struct {
	SourceId int `json:"source_id"`
	TargetId int `json:"target_id"`
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Delete struct {
	Id int `json:"target_id"`
}

type ResponseList struct {
	Items []*User
}
