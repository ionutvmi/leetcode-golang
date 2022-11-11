import "strings"

func romanToInt(s string) int {
    chars := strings.Split(s, "")
    
    result := 0
    
    symbolsMap := map[string]int{
        "I": 1,
        "IV": 4,
        "V": 5,
        "IX": 9,
        "X": 10,
        "XL": 40,
        "L": 50,
        "XC": 90,
        "C": 100,
        "CD": 400,
        "D": 500,
        "CM": 900,
        "M": 1000,
    }

    complexSymbols := []string{"IV", "IX", "XL", "XC", "CD", "CM"}
    size := len(chars)
    index := 0
    
    for index < size {
        if index < size - 1 {
            candidate := chars[index] + chars[index+1]
            if contains(complexSymbols, candidate) {
                result += symbolsMap[candidate]
                index += 2
                continue
            }
        }
        candidate := chars[index]
        result += symbolsMap[candidate]
        index += 1
    }
    
    return result
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
