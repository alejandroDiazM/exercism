"""Solution to Ellen's Alien Game exercise."""

class Alien:
    """Create an Alien object with location x_coordinate and y_coordinate.

    Attributes
    ----------
    (class)total_aliens_created: int
    x_coordinate: int - Position on the x-axis.
    y_coordinate: int - Position on the y-axis.
    health: int - Amount of health points.

    Methods
    -------
    hit(): Decrement Alien health by one point.
    is_alive(): Return a boolean for if Alien is alive (if health is > 0).
    teleport(new_x_coordinate, new_y_coordinate): Move Alien object to new coordinates.
    collision_detection(other): Implementation TBD.
    """

    total_aliens_created = 0
    health = 3
    
    def __init__(self, x_coordinate, y_coordinate):
        Alien.total_aliens_created += 1
        self.x_coordinate = x_coordinate
        self.y_coordinate = y_coordinate
    
    def hit(self):
        self.health -= 1
    
    def is_alive(self):
        return self.health > 0
    
    def teleport(self, new_x_coordinate, new_y_coordinate):
        self.x_coordinate = new_x_coordinate
        self.y_coordinate = new_y_coordinate
    
    def collision_detection(self, other):
        pass

    def new_aliens_collection(positions):
        """Creates a list of Alien instances from a list of coordinate tuples.
    
        :param positions: list - a list of tuples of (x, y) coordinates.
        :return: list - a list of Alien objects.
        """
        return [Alien(position[0], position[1]) for position in positions]