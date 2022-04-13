from audioop import reverse


def latest(scores):
    return scores[-1]


def personal_best(scores):
    return max(scores)


def personal_top_three(scores):
    scores_copy = list(scores)
    scores_copy.sort()
    highest_three = scores_copy[-3:]
    highest_three.reverse()
    return highest_three
