package par

import (
	"encoding/json"
	"testing"
)

var data = Response{}

func init() {
	const n = 3

	data.Bids = make([]Bid, 0, n)

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

func BenchmarkResponse(b *testing.B) {
	var size int64

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err := json.Marshal(&data)
		if err != nil {
			b.Error(err)
		}

		size = int64(len(res))
	}

	b.SetBytes(size)
}
