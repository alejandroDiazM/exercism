def response(hey_bob: str):
    hey_bob = hey_bob.strip()

    yelling = hey_bob.isupper()
    question = hey_bob.endswith("?")

    if hey_bob == "":
        return "Fine. Be that way!"
    elif yelling and question:
        return "Calm down, I know what I'm doing!"
    elif yelling:
        return "Whoa, chill out!"
    elif question:
        return "Sure."
    else:
        return "Whatever."
