package hackernews

type Item struct {
	Id          int    `json:"id"`
	Deleted     string `json:"deleted"`
	Text        string `json:"text"`
	Dead        string `json:"dead"`
	Parent      int    `json:"parent"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	Kids        []int  `json:"kids"`
	Parts       []int  `json:"parts"`
}

type User struct {
	About     string `json:"about"`
	Created   int    `json:"created"`
	Delay     int    `json:"delay"`
	Id        string `json:"id"`
	Karma     int    `json:"karma"`
	Submitted []int  `json:"submitted"`
}
