def romanToInt(s: str) -> int:
    sum = 0
    values = {
        'I': 1,
        'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
    }

    for i, ch in enumerate(s):
        cur_val =  values[ch]
        next_val = 0
        if i+1 < len(s):
            next_val = values[s[i+1]]
        if cur_val < next_val:
            sum -= cur_val
        else:
            sum += cur_val

    return sum

if __name__ == "__main__":
    print(romanToInt("MCMXCIV"))