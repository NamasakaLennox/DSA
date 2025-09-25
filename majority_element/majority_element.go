func majorityElement(nums []int) int {
    result := 0
    majority := 0

    for _, val := range(nums) {
        if (majority == 0){
            result = val
        }
        if (val == result) {
            majority++
        } else {
            majority--
        }
    }
    return result
}
