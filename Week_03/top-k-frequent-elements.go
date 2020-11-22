func topKFrequent(nums []int, k int) []int {
	var result []int
	hashMap := make(map[int]int, 0)

	for _, v := range nums {
		hashMap[v]++
	}

	pairs := make(Pairs, len(hashMap))

	i := 0
	for key, v := range hashMap {
		pairs[i] = Pair{
			Key:   key,
			Value: v,
		}
		i++
	}
	sort.Sort(pairs)

	for _, p := range pairs[:k] {
		result = append(result, p.Key)
	}
	return result
}

type Pair struct {
	Key   int
	Value int
}

type Pairs []Pair

func (ps Pairs) Len() int           { return len(ps) }
func (ps Pairs) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }
func (ps Pairs) Less(i, j int) bool { return ps[i].Value > ps[j].Value }