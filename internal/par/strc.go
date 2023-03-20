package par

type Bid struct {
	Title     string   `json:"title"`
	Text      string   `json:"text"`
	ClickUrl  string   `json:"click_url"`
	IconUrl   string   `json:"icon_url"`
	ImageUrl  string   `json:"image_url"`
	RequestId string   `json:"request_id"`
	Pixels    []string `json:"pixels"`
}

type Response struct {
	Bids []Bid `json:"bids"`
}
