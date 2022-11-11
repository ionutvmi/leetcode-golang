import "strings"


func intToRoman(num int) string {
    symbolsDict := map[int]string{
        1000: "M",
        900: "CM",
        500: "D",
        400: "CD",
        100: "C",
        90: "XC",
        50: "L",
        40: "XL",
        10: "X",
        9: "IX",
        5: "V",
        4: "IV",
        1: "I",
    }
    orderedParts := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
    
    result := ""
    
    for _, part := range orderedParts {
        for num >= part {
            result += symbolsDict[part]
            num -= part
        }
    }
    
    
    return result
}


func intToRoman_firstAttempt(num int) string {
    result := ""
    reminder := num
    
    if reminder >= 1000 {
        result += strings.Repeat("M", int(num / 1000))
        reminder = reminder % 1000
    }

    if reminder >= 900 {
        result += "CM"
        
        reminder = reminder % 900
    }

    
    if reminder >= 500 {
        result += "D"
        
        reminder = reminder % 500
    }
    
    if reminder >= 400 {
        result += "CD"
        
        reminder = reminder % 400
    }
    
    if reminder >= 100 {
        result += strings.Repeat("C", int(reminder / 100))
        reminder = reminder % 100
    }
    
    if reminder >= 90 {
        result += "XC"
        
        reminder = reminder % 90
    }
    
    if reminder >= 50 {
        result += "L"
        
        reminder = reminder % 50
    }
    
    if reminder >= 40 {
        result += "XL"
        
        reminder = reminder % 40
    }
    
    if reminder >= 10 {
        result += strings.Repeat("X", int(reminder / 10))
        reminder = reminder % 10
    }
    
    if reminder == 9 {
        result += "IX"
        reminder -= 9
    }
    
    if reminder >= 5 {
        result += "V"
        reminder = reminder % 5
    }
    
    if reminder == 4 {
        result += "IV"
        reminder -= 4
    }

    result += strings.Repeat("I", reminder)

    return result
}

