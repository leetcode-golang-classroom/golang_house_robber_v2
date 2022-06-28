# golang_house_robber_v2

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are **arranged in a circle.** That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have a security system connected, and **it will automatically contact the police if two adjacent houses were broken into on the same night**.

Given an integer array `nums` representing the amount of money of each house, return *the maximum amount of money you can rob tonight **without alerting the police***.

## Examples

**Example 1:**

```
Input: nums = [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.

```

**Example 2:**

```
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.

```

**Example 3:**

```
Input: nums = [1,2,3]
Output: 3

```

**Constraints:**

- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 1000`

## 解析

跟 [**198. House Robber**](https://www.notion.so/198-House-Robber-82bdf8d17ad441beaa73755bd899ab7e) 類似

題目給定一個陣列 nums ， 每個 nums[i] 該戶有的錢數目

相鄰的兩間如果連續拿錢就會啟動警報

而這題與[**198. House Robber**](https://www.notion.so/198-House-Robber-82bdf8d17ad441beaa73755bd899ab7e) 不同的是， 這題的第1間與最後一間是相鄰的是一個環狀結構

要求實做出一個演算法在不啟動警報的情況下 拿出最多錢

因為第0間與最後一間是相鄰的所以不能直接用上一題的解法

透過下圖

![](https://i.imgur.com/FzW93yP.png)

可以發現只要把所有用戶分成第 0 到 len(nums)-2 家與 第 1 到 len(nums)-1 家

分別去考慮這樣就可以把第0間 跟最後一間分開可以使用 [**198. House Robber**](https://www.notion.so/198-House-Robber-82bdf8d17ad441beaa73755bd899ab7e) 的解法

分別考慮 0 到 len(nums)-2 與 1 到 len(nums) -1 然後取最大值

就是結果

由於上一個 [**198. House Robber**](https://www.notion.so/198-House-Robber-82bdf8d17ad441beaa73755bd899ab7e) 的解為 O(n)

所以這題的時間複雜度也是 O(n)

## 程式碼
```go
package sol

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func rob(nums []int) int {
	total := len(nums)

	var rob_with_range = func(start, end int) int {
		prevTwo := 0
		prevOne := nums[end]
		totalRob := 0
		for robIdx := end - 1; robIdx >= start; robIdx-- {
			totalRob = max(nums[robIdx]+prevTwo, prevOne)
			prevTwo = prevOne
			prevOne = totalRob
		}
		return max(prevOne, prevTwo)
	}
	if total == 1 {
		return nums[0]
	}

	return max(rob_with_range(0, total-2), rob_with_range(1, total-1))
}

```
## 困難點

1. 要看出可以把環狀結構拆分成兩個部份

## Solve Point

- [x]  把 index 分成 前 n-2 個 與後 n-2 個 分別去呼叫 [**198. House Robber**](https://www.notion.so/198-House-Robber-82bdf8d17ad441beaa73755bd899ab7e) 的解法
- [x]  然後把兩個結果