class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        k = 1
        i = 1
        j = 1
        while j < len(nums):
            if nums[i - 1] != nums[j]:
                nums[i] = nums[j]
                k += 1
                i += 1
            j += 1
        return k
