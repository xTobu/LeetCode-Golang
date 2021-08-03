# LeetCode-Golang
> jotting

![Author](https://img.shields.io/badge/Author-Junxiang-yellow.svg)

---
## Problems
| No    | #     | Source                                           | Solution                  | Difficulty | 
| :---: | :---: | :----------------------------------------------: | :-----------------------: | :--------: |
| 1     | 1     | [Two Sum][#1]                                    | [Link](/algorithms/0001.Two-Sum)    | Easy       | 
| 2     | 21   | [Merge Two Sorted Lists][#21]                       | [Link](/algorithms/0021.Merge-Two-Sorted-Lists)  | Easy       |
| 3     | 121   | [Best Time to Buy and Sell Stock][#121]                        | [Link](/algorithms/0121.Best-Time-to-Buy-and-Sell-Stock)  | Easy       |
| 4     | 136   | [Single Number][#136]                  | [Link](/algorithms/0136.Single-Number)  | Easy       |
| 5     | 169   | [Majority Element][#169]                        | [Link](/algorithms/0169.Majority-Element)  | Easy       | 
| 6     | 189   | [Rotate Array][#189]        | [Link](/algorithms/0189.Rotate-Array)  | Medium     | 
| 7     | 206   | [Reverse Linked List][#206]                      | [Link](/algorithms/0206.Reverse-Linked-List)  | Easy     | 
| 8     | 217   | [Contains Duplicate][#217]                         | [Link](/algorithms/0217.ContainsDuplicate)  | Easy       | 
| 9     | 242   | [Valid Anagram][#242]                      | [Link](/algorithms/0242.Valid-Anagram)  | Easy     | 
| 10    | 283   | [Move Zeroes][#283]                   | [Link](/algorithms/0283.Move-Zeroes)  | Easy       | 
| 11...    |    | 解題忘了填，待補                |   |        | 

<!-- 參考 超連結 Source -->
[#1]: https://leetcode.com/problems/two-sum/description/
[#21]:https://leetcode.com/problems/merge-two-sorted-lists/
[#121]:https://leetcode.com/problems/best-time-to-buy-and-sell-stock/  
[#136]:https://leetcode.com/problems/single-number/
[#169]:https://leetcode.com/problems/majority-element/
[#189]:https://leetcode.com/problems/rotate-array/
[#206]:https://leetcode.com/problems/reverse-linked-list/
[#217]:https://leetcode.com/problems/contains-duplicate/
[#242]:https://leetcode.com/problems/valid-anagram/
[#283]:https://leetcode.com/problems/move-zeroes/


## Testing

### Run Test

```bash
bash ./test.sh
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
