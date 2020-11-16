func twoSum(nums []int, target int) []int {
    hashTable := map[int]int{}

    for i, v := range nums {
        if t, ok := hashTable[target - v]; ok {
            return []int{t, i}
        }
        hashTable[v] = i
    }
    return nil
}