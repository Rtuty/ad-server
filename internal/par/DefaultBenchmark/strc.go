//go:generate easyjson -no_std_marshalers esjsn.go
package par

//easyjson:json
type (
	Bid struct {
		Title     string   `json:"title"`
		Text      string   `json:"text"`
		ClickUrl  string   `json:"click_url"`
		IconUrl   string   `json:"icon_url"`
		ImageUrl  string   `json:"image_url"`
		RequestId int      `json:"request_id"`
		Pixels    []string `json:"pixels"`
	}

	Response struct {
		Bids []Bid `json:"bids"`
	}
)
