func permute(nums []int) [][]int {
	result := [][]int{}
	visited := make(map[int]bool)

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for _, n := range nums {
			if visited[n] {
				continue
			}

			visited[n] = true
			path = append(path, n)
			dfs(path)
			path = path[:len(path)-1]
			visited[n] = false
		}
	}

	dfs([]int{})

	return result
}