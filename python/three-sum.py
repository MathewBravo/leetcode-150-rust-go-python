from typing import List

def threeSum(nums: List[int]) -> List[List[int]]:
    result = []
    def quick_sort(nums) -> List[int]:
        if len(nums) <= 1:
            return nums
        pivot = nums[0]
        less, greater = [], []
        for num in nums[1:]:
            if num < pivot:
                less.append(num)
            else:
                greater.append(num)
        
        less = quick_sort(less)
        greater = quick_sort(greater)
        return less + [pivot] + greater

    nums = quick_sort(nums)

    for i, num in enumerate(nums):
        if i > 0 and nums[i-1] == num:
            continue
        
        left, right = i+1, len(nums)-1
        while left < right:
            sumTotal = num + nums[left] + nums[right]
            if sumTotal > 0:
                right -= 1
            elif sumTotal < 0:
                left += 1
            else:
                result.append([num, nums[left], nums[right]])
                left += 1
                right -= 1
                while left < right and nums[left-1] == nums[left]:
                    left += 1
                while right > left and nums[right+1] == nums[right]:
                    right -= 1

    return result

    
if __name__ == "__main__":
    res = threeSum([-1,0,1,2,-1,-4])
    print(res)

