from typing import List

def twoSum(numbers: List[int], target: int) -> List[int]:
    seen_nums = {}
    for index, num in enumerate(numbers):
        seen_nums[num] = index

    for num in seen_nums:
        if(target-num in seen_nums):
            return [seen_nums[num]+1, seen_nums[target-num]+1]
    return [-1,-1]

if __name__ == "__main__":
    print(twoSum([2,7,11,15], 9))