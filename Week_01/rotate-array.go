func rotate(nums []int, k int)  {
    size := len(nums)
    count := 0
    k = k % len(nums)
    // position
    prev := 0
    prevValue := 0
    next := 0
    nextValue := 0

    for i, _ := range nums {
        prev = i 
        prevValue = nums[i]
        for count < size {
            next = (prev+k) % size
            nextValue = nums[next]
            nums[next] = prevValue
            count++

            prev = next
            prevValue = nextValue
            if prev == i {
                break
            }
        } 
    }
}