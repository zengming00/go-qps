package main

import (
	"math/rand"
	"net/http"
	"time"

	mqps "github.com/zengming00/go-qps"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Statistics every second, a total of 3600 data
	qps := mqps.NewQP(time.Second, 3600)
	for i := 0; i < 5; i++ {
		go func() {
			for {
				// Call Count() on every goroutine that needs statistics
				qps.Count()
				time.Sleep(time.Millisecond * time.Duration((50 + rand.Intn(500))))
			}
		}()
	}
	// Add a route to get HTML, for example /qps
	http.HandleFunc("/qps", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// Get the raw HTML, you can gzip it
		s, err := qps.Show()
		if err != nil {
			panic(err)
		}
		w.Write([]byte(s))
	})
	// Add a route to get json report(for automatic refresh), The name is the same as getting the HTML routing, but you need to add the '_json' suffix
	http.HandleFunc("/qps_json", func(w http.ResponseWriter, r *http.Request) {
		// Get the json report
		bts, err := qps.GetJson()
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bts)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
