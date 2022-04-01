def response(hey_bob):
    hey_bob = str(hey_bob)
    
    if hey_bob.endswith("?") and not hey_bob.isupper():
        return "Sure."
    elif hey_bob.isupper():
        return "Whoa, chill out!"
    elif hey_bob.isupper() and "?" in hey_bob:
        return "Calm down, I know what I'm doing!"
    elif not hey_bob:
        return "Fine. Be that way!"
    else:
        return "Whatever."