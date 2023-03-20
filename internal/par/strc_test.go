package par

var data = Response{}

func init() {
	const n = 3

	data.Bids = make([]Bid, 8, n)

	for i := 0; i < n; i++ {
		data.Bids = append(data.Bids,
			Bid{
				Title:     "Base title",
				Text:      "Base test",
				ClickUrl:  "http://example.com",
				IconUrl:   "http://example.com/icon.png",
				ImageUrl:  "http://example.com/image.png",
				RequestId: i,
				Pixels: []string{
					"http://example.com/pixel1",
					"http://example.com/pixel1",
					"http://example.com/pixel1",
				},
			})
	}
}
