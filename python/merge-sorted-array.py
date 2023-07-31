from typing import List

def merge(nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        # initialize pointers for nums1, nums2, and the end of the merged array
    p1, p2, p3 = m - 1, n - 1, m + n - 1
    
    # merge the arrays in reverse order
    while p1 >= 0 and p2 >= 0:
        if nums1[p1] >= nums2[p2]:
            nums1[p3] = nums1[p1]
            p1 -= 1
        else:
            nums1[p3] = nums2[p2]
            p2 -= 1
        p3 -= 1
    
    # if there are remaining elements in nums2, copy them to nums1
    while p2 >= 0:
        nums1[p3] = nums2[p2]
        p2 -= 1
        p3 -= 1




if __name__ == "__main__":
    merge([1,2,3,0,0,0], 3, [2,5,6], 3)