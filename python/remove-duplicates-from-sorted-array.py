from typing import List

def removeDuplicates(nums: List[int]) -> int:
    if len(nums) <= 1:
        return 1
    uniq_i = 0
    for num in nums:
        if num != nums[uniq_i]:
            uniq_i += 1
            nums[uniq_i] = num
    return uniq_i + 1

if __name__ == "__main__":
    print(removeDuplicates([0,0,1,1,1,2,2,3,3,4]))