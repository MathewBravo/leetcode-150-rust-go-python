from typing import List

def productExceptSelf(nums: List[int]) -> List[int]:
    length = len(nums)
    prefix = [1] * length
    suffix = [1] * length
    
    for i in range(1, length):
        prefix[i] = prefix[i - 1] * nums[i - 1]
    for i in range(length - 2, -1, -1):
        suffix[i] = suffix[i + 1] * nums[i + 1]
    
    result = [prefix[i] * suffix[i] for i in range(length)]
    
    return result

if __name__ == "__main__":
    print(productExceptSelf([1,2,3,4]))