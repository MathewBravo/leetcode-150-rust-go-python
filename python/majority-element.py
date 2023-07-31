from typing import List
def majorityElement(nums: List[int]) -> int:
    el_count = {}
    for num in nums:
        if num in el_count:
            el_count[num] = el_count[num] + 1
        else:
            el_count[num] = 1
    return max(el_count, key=el_count.get)

if __name__ == "__main__":
    print(majorityElement([3,2,3]))