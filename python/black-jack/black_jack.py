"""Functions to help play and score a game of blackjack.

How to play blackjack:    https://bicyclecards.com/how-to-play/blackjack/
"Standard" playing cards: https://en.wikipedia.org/wiki/Standard_52-card_deck
"""


def value_of_card(card):
    """Determine the scoring value of a card.

    :param card: str - given card.
    :return: int - value of a given card. 'J', 'Q', 'K' = 10; 'A' = 1; numerical value otherwise.
    """

    if card == "2":
        return 2
    if card == "3":
        return 3
    if card == "4":
        return 4
    if card == "5":
        return 5
    if card == "6":
        return 6
    if card == "7":
        return 7
    if card == "8":
        return 8
    if card == "9":
        return 9
    elif card == "J" or card == "Q" or card == "K" or card == "10":
        return 10
    elif card == "A":
        return 1


def higher_card(card_one, card_two):
    """Determine which card has a higher value in the hand.

    :param card_one, card_two: str - cards dealt. 'J', 'Q', 'K' = 10; 'A' = 1; numerical value otherwise.
    :return: higher value card - str. Tuple of both cards if they are of equal value.
    """

    if value_of_card(card_one) > value_of_card(card_two):
        return card_one
    elif value_of_card(card_one) < value_of_card(card_two):
        return card_two
    elif value_of_card(card_one) == value_of_card(card_two):
        return card_one, card_two


def value_of_ace(card_one, card_two):
    """Calculate the most advantageous value for the ace card.

    :param card_one, card_two: str - card dealt. 'J', 'Q', 'K' = 10;
           'A' = 11 (if already in hand); numerical value otherwise.

    :return: int - value of the upcoming ace card (either 1 or 11).
    """
    sumOfValues = value_of_card(card_one) + value_of_card(card_two)
    aceInHand = card_one == "A" or card_two == "A"
    noAceInHand = card_one != "A" and card_two != "A"

    if sumOfValues <= 10 and noAceInHand:
        return 11
    elif sumOfValues > 10 or aceInHand:
        return 1
    

def is_blackjack(card_one, card_two):
    """Determine if the hand is a 'natural' or 'blackjack'.

    :param card_one, card_two: str - cards dealt. 'J', 'Q', 'K' = 10; 'A' = 11; numerical value otherwise.
    :return: bool - if the hand is a blackjack (two cards worth 21).
    """
    hand = card_one, card_two
    valueOfTenInHand = "K" in hand or "Q" in hand or "J" in hand or "10" in hand

    if "A" in hand and valueOfTenInHand:
        return True
    else:
        return False
    


def can_split_pairs(card_one, card_two):
    """Determine if a player can split their hand into two hands.

    :param card_one, card_two: str - cards dealt.
    :return: bool - if the hand can be split into two pairs (i.e. cards are of the same value).
    """
    if value_of_card(card_one) == value_of_card(card_two):
        return True
    else:
        return False
    


def can_double_down(card_one, card_two):
    """Determine if a blackjack player can place a double down bet.

    :param card_one, card_two: str - first and second cards in hand.
    :return: bool - if the hand can be doubled down (i.e. totals 9, 10 or 11 points).
    """
    doubleDownCond = value_of_card(card_one) + value_of_card(card_two) == 9 or\
    value_of_card(card_one) + value_of_card(card_two) == 10 or\
    value_of_card(card_one) + value_of_card(card_two) == 11

    if doubleDownCond:
        return True
    else:
        return False
    
