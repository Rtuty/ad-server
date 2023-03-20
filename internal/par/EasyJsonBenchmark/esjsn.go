package par

import (
	"encoding/json"
	"github.com/mailru/easyjson"
	par "modules/internal/par/DefaultBenchmark"
	"testing"
)

package par

import (
"encoding/json"
"testing"
)

var data = par.Response{}

func init() {
	const n = 3

	data.Bids = make([]par.Bid, 0, n)

	for i := 0; i < n; i++ {
		data.Bids = append(data.Bids,
			par.Bid{
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

func BenchmarkResponseEasyJson(b *testing.B) {
	var size int64

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err := easyjson.Marshal(&data)
		if err != nil {
			b.Error(err)
		}

		size = int64(len(res))
	}

	b.SetBytes(size)
}
