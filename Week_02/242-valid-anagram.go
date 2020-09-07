func isAnagram(s string, t string) bool {
    tmpMap := make(map[byte]int)

    lenS := len(s)
    lenT := len(t)

    if lenS != lenT {
        return false
    }

    for i:=0; i<lenS; i++ {
        tmpMap[s[i]]++
    }

    for i:=0; i<lenT; i++ {
        _, exists := tmpMap[t[i]]
        if !exists {
            return false
        }
        tmpMap[t[i]]--
    }

    for _, v := range tmpMap {
        if v != 0 {
            return false
        }
    }

    return true
}