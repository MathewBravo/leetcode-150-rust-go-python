from typing import List

def maxProfit(prices: List[int]) -> int:
        min_price = float('inf')
        max_profit = 0
        for i in range(len(prices)):
            if prices[i] < min_price:
                min_price = prices[i]
            elif prices[i] - min_price > max_profit:
                max_profit = prices[i] - min_price
                
        return max_profit

        
if __name__ == "__main__":
    print(maxProfit([7,1,5,3,6,4]))