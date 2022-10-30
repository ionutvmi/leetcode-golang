import "strings"

func lengthOfLongestSubstring_initialAttempt(s string) int {
    result := ""
    chars := strings.Split(s, "")
    size := len(chars)
   
    for index, c := range chars {
        
        candidate := c

        for i := index + 1; i < size; i++ {
            newChar := chars[i]

            if strings.Contains(candidate, newChar) {
                break
            }
            
            candidate += newChar
        }
        
        if len(candidate) > len(result) {
            result = candidate
        }
        
    }
    
    return len(result)
}

func MaxOf(vars ...int) int {
    max := vars[0]

    for _, i := range vars {
        if max < i {
            max = i
        }
    }

    return max
}

func lengthOfLongestSubstring(s string) int {
    result := 0
    leftIndex := 0
    seenMap := map[string]int{}
    
    chars := strings.Split(s, "")
    
    for rightIndex, char := range chars {
        
        seenIndex, ok := seenMap[char]
        
        if !ok {
            result = MaxOf(result, rightIndex - leftIndex + 1)
        } else {
            if seenIndex < leftIndex {
                result = MaxOf(result, rightIndex - leftIndex + 1)
            } else {
                leftIndex = seenIndex + 1
            }
        }
        
        seenMap[char] = rightIndex
        
    }
    
    
    return result
}
