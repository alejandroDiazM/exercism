class Matrix:
    def __init__(self, matrix_string: str):
        self.matrix_list = matrix_string.split("\n")
        self.matrix_list = [[int(i) for i in row.split()] for row in self.matrix_list]

    def row(self, index):
        return self.matrix_list[index-1]

    def column(self, index):
        return [row[index-1] for row in self.matrix_list]
