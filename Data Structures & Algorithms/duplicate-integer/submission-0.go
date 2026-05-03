func hasDuplicate(nums []int) bool {
    numMap := make(map[int]bool)
    for _, val := range (nums) {
        if _,ok := numMap[val]; !ok {
            numMap[val] = false
        }else{
            return true
        }
    }
    return false
}