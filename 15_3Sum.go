import "sort"
import "strconv"

func threeSum(nums []int) [][]int {
    resultsMap := map[string]bool{}

    numsMap := map[int]int{}
    
    for i, n := range nums {
        numsMap[n] = i
    }
    
    result := [][]int{}
    
    for i, firstNumber := range nums {
        
        for j, secondNumber := range nums {
            
            if i == j {
                continue
            }
            
            thirdNumber := 0 - (firstNumber + secondNumber)
            
            k, ok := numsMap[thirdNumber]  // O(1)
            
            if ok && i != k && j != k {
                hash := getHash(nums[i], nums[j], nums[k])
                if resultsMap[hash] {
                    continue
                }                
                
                result = append(result, []int{nums[i], nums[j], nums[k]})
                resultsMap[hash] = true
            }
        }
        
    }
    
    return result
}


func getHash(args... int) string {
    sort.Ints(args)

    result := ""

    for _, v := range args {
        result += strconv.Itoa(v)
    }
    
    return result
}
