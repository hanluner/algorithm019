func groupAnagrams(strs []string) [][]string {
    var res [][]string

    hashMap := make(map[string]string)
    for _, s := range strs {
        key := sortStr(s)
        va, ok := hashMap[key]
        if ok {
            hashMap[key] = va + "," + s
        } else {
            hashMap[key] = s
        }
    }
    for _, v := range hashMap {
        tmpArr := strings.Split(v, ",")
        res = append(res, tmpArr)
    }

    return res
}

func sortStr(s string) string {
    var sArr []string

    for _, v := range s {
        sArr = append(sArr, string(v))
    }
    sort.Strings(sArr)

    return strings.Join(sArr, "")
}