`go run latency.go`

Makes 10 total trials for two files using range and non-range requests.

from asia:
```
544.00 ms (small file)
354.70 ms (small file range)
504.10 ms (big file)
556.10 ms (big file range begin)
687.00 ms (big file range end)
```

from us west:
```
178.90 ms (small file)
168.00 ms (small file range)
177.60 ms (big file)
206.00 ms (big file range begin)
357.60 ms (big file range end)
```
