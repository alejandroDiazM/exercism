resistor_colors = [
        "black",
        "brown",
        "red",
        "orange",
        "yellow",
        "green",
        "blue",
        "violet",
        "grey",
        "white",
    ]

def color_code(color: str):
    lower_color = color.lower()
    return resistor_colors.index(lower_color)


def colors():
    return resistor_colors
