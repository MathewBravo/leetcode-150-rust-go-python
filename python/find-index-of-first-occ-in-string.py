def strStr(haystack: str, needle: str) -> int:
    if haystack == needle:
        return 0
    n_len = len(needle)
    for i, ch in enumerate(haystack):
        if haystack[i:i + n_len] == needle:
            return i
    return -1

if __name__ == "__main__":
    print(strStr("sadbutsad", "sad"))
    print(strStr("leetcode", "leeto"))