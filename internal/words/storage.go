package words

type Word struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Translation string `json:"translation"`
}
