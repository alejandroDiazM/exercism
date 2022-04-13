def abbreviate(words):
    words = words.replace("-", " ").replace("_", "").split()
    abbr = "".join([word[0].upper() for word in words])
    return abbr