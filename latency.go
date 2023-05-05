package main

import (
	"fmt"
	"net/http"
	"time"
)


func main() {
	client := http.Client{}

	measureTime := func(url string, isRange bool, byteToRequest uint64) float64 {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			panic(err)
		}

		rng := ""
		if (isRange) {
			rng = fmt.Sprintf("bytes=%d-%d",(byteToRequest), (byteToRequest))
			req.Header.Set("Range", rng)
		}

		startTime := time.Now()
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		duration := time.Since(startTime)
		defer resp.Body.Close()

		return float64(duration.Milliseconds())
	}

	TRIALS := 10

	small_file := "https://r2-public.protomaps.com/protomaps-sample-datasets/1byte.txt"
	big_file := "https://r2-public.protomaps.com/protomaps-sample-datasets/terrarium-z12.pmtiles"

	small_file_norange := 0.0
	small_file_range := 0.0
	big_file_norange := 0.0
	big_file_rangebegin := 0.0
	big_file_rangeend := 0.0

	for i := 0; i < TRIALS; i++ {
		small_file_norange += measureTime(small_file, false, 0)
		small_file_range += measureTime(small_file, true, 0)
		big_file_norange += measureTime(big_file, false, 0)
		big_file_rangebegin += measureTime(big_file, true, 0)
		big_file_rangeend += measureTime(big_file, true, 159822873634)
	}

	fmt.Printf("%.2f ms (small file)\n", small_file_norange / float64(TRIALS))
	fmt.Printf("%.2f ms (small file range)\n", small_file_range / float64(TRIALS))
	fmt.Printf("%.2f ms (big file)\n", big_file_norange / float64(TRIALS))
	fmt.Printf("%.2f ms (big file range begin)\n", big_file_rangebegin / float64(TRIALS))
	fmt.Printf("%.2f ms (big file range end)\n", big_file_rangeend / float64(TRIALS))
}