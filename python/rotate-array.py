from typing import List
def rotate(nums: List[int], k: int) -> None:
    k = k % len(nums)

    def reverse_in_place(nums, left, right):
        while left < right:
            nums[left], nums[right] = nums[right], nums[left]
            left += 1
            right -= 1

    reverse_in_place(nums, 0, len(nums) - 1)
    reverse_in_place(nums, 0, k-1)
    reverse_in_place(nums, k, len(nums)-1)
    return nums

if __name__ == "__main__":
    print(rotate([1,2,3,4,5,6,7], 3))