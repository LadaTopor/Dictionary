package reports

type Reports struct {
	Id          int    `json:"id"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}
