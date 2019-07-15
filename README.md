# LeetCode-Golang
LeetCode with Golang

---

## Testing

### Run Test

```bash
bash ./test.bash
```

### Display details
```bash
# Function
go tool cover -func=coverage.out

# HTML
go tool cover -html=coverage.out
```

### Benchmark
```bash
# -bench=.  -> run benchmark
# -run=none -> no testing
go test -v -bench=. -run=none .
```