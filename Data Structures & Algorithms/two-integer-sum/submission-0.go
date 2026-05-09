func twoSum(nums []int, target int) []int {
    numMap :=  make(map[int]int)
    for i,val := range nums {
        numMap[val] = i
    }
    for i,val := range nums {
        j,ok := numMap[target-val]
        if ok && i != j {
            return []int{i,j}
        }
    }
    return []int{}
}