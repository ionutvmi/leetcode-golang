import "strings"
import "strconv"

func isPalindrome(x int) bool {
    
    if x < 0 {
        return false
    }
    
    chars := strings.Split(strconv.Itoa(x), "")
    size := len(chars)
    half := size / 2


    for i := 0; i < half; i++ {
        if chars[i] != chars[size-i-1] {
            return false
        }
    }
    
    return true
    
}
