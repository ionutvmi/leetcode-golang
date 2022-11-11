import "strings"

func longestCommonPrefix(strs []string) string {
    
    prefix := ""
    firstWord := strs[0]
    
    if len(firstWord) == 0 {
        return prefix
    }
    
    chars := strings.Split(firstWord, "")
    for _, c := range chars {
        candidate := prefix + c
        
        for _, s := range strs {
            if !strings.HasPrefix(s, candidate) {
                return prefix
            }
        }

        prefix = candidate
    }
    
    return prefix
}
