def is_isogram(string: str):
    string = string.lower()
    string = string.replace("-", "")
    string = string.replace(" ", "")
    return len(set(string)) == len(string)
