func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sMap := make(map[rune]int)
    for _, v := range s {
        sMap[v]++
    }

    for _, v := range t {
        value, ok := sMap[v]
        if !ok {
            return false
        }
        if value == 1 {
            delete(sMap, v)
        } else {
            sMap[v]--
        }
    }
    if len(sMap) == 0 {
        return true
    }
    return false
}