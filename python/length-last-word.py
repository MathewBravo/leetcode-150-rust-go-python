def lengthOfLastWord(s: str) -> int:
    new_string = " ".join(s.split())
    split_string = new_string.split(" ")
    return len(split_string[-1])

if __name__ == "__main__":
    print(lengthOfLastWord("Hello World"))