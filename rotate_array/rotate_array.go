func reverse(arr []int, start int, end int) {
    for start < end {
        arr[start], arr[end] = arr[end], arr[start]
        start++
        end--
    }
}

func rotate(nums []int, k int)  {
    arr_len := len(nums)
    k = k % arr_len
    if (arr_len <= 1 || k == 0) {
        return
    }
    reverse(nums, 0, arr_len - 1)
    reverse(nums, 0, k - 1)
    reverse(nums, k, arr_len - 1)
}
