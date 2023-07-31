from typing import List

def removeElement(nums: List[int], val: int) -> int:
    slow = 0
    for num in nums:
        if num != val:
            nums[slow] = num
            slow += 1 
        
    return slow

if __name__ == "__main__":
    print(removeElement([3,2,2,3], 3))