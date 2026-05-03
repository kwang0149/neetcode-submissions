func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    mapA := make(map[byte]int)
    for i := 0;i < len(s); i ++ {
        mapA[s[i]]++
        mapA[t[i]]--
    }
    for _,val := range mapA {
        if val != 0 {
            return false
        }
    }
    return true
}