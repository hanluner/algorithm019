func generateParenthesis(n int) []string {
	res := new([]string)
	backTree(n, n, n, "", res)
	return *res
}

func backTree(left int, right, n int, path string, res *[]string) {
	// terminator
	if right == 0 {
		*res = append(*res, path)
		return
	}

	if left > 0 {
		backTree(left-1, right, n, path+"(", res)
	}

	if left < right {
		backTree(left, right-1, n, path+")", res)
	}

	return
}