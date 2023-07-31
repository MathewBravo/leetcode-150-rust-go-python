from typing import List
def trap(height: List[int]) -> int:
    left, right = 0, len(height)-1
    max_left, max_right, sum = height[left], height[right],0 
    while left < right:
        test1 = height[left]
        test2 = height[right]
        if max_left < max_right:
            left += 1
            water = max_left - height[left]
            if water > 0:
                sum += water
            max_left = max(max_left, height[left])
        else:
            right -= 1
            water = max_right - height[right]
            if water > 0:
                sum += water
            max_right = max(max_right, height[right])
    return sum

if __name__ == "__main__":
    print(trap([0,1,0,2,1,0,1,3,2,1,2,1]))