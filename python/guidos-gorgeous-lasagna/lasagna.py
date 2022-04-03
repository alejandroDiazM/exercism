"""Functions used in preparing Guido's gorgeous lasagna.

Learn about Guido, the creator of the Python language: https://en.wikipedia.org/wiki/Guido_van_Rossum
"""

# TODO: define the 'EXPECTED_BAKE_TIME' constant
# TODO: consider defining the 'PREPARATION_TIME' constant
#       equal to the time it takes to prepare a single layer
from unittest import result


EXPECTED_BAKE_TIME = 40
PREPARATION_TIME = 2

# TODO: define the 'bake_time_remaining()' function
def bake_time_remaining(elapsed_bake_time):
    """Calculate the bake time remaining.

    :param elapsed_bake_time: int - baking time already elapsed.
    :return: int - remaining bake time derived from 'EXPECTED_BAKE_TIME'.

    Function that takes the actual minutes the lasagna has been in the oven as
    an argument and returns how many minutes the lasagna still needs to bake
    based on the `EXPECTED_BAKE_TIME`.
    """
    result = EXPECTED_BAKE_TIME - elapsed_bake_time
    return result

# TODO: define the 'preparation_time_in_minutes()' function
#       and consider using 'PREPARATION_TIME' here
def preparation_time_in_minutes(number_of_layers):
    """Calculate the total bake time needed based on the number of layers.    
    """
    result = PREPARATION_TIME * number_of_layers
    return result

# TODO: define the 'elapsed_time_in_minutes()' function
def elapsed_time_in_minutes(number_of_layers, elapsed_bake_time):
    """Calculate the total elapsed cooking time (prep + bake) in minutes
    """
    layersTime = number_of_layers * PREPARATION_TIME
    result = layersTime + elapsed_bake_time
    return result