func lengthOfLongestSubstring(s string) int {
    res := 0
    for i := 0; i < len(s); i++ {
        set := make(map[byte]bool)
        length := 0

        for j := i; j < len(s); j++ {
            if set[s[j]] {
                break
            }

            set[s[j]] = true
            length++
        }

        if res < length {
            res = length
        }
    }

    return res
}