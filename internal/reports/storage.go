package reports

type Reports struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}
