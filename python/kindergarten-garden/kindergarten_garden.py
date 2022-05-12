full_student_list = ['Alice', 'Bob', 'Charlie', 'David', 'Eve', 'Fred', 'Ginny', 'Harriet', 'Ileana', 'Joseph', 'Kincaid', 'Larry']

plants = {'C':'Clover', 'G':'Grass', 'R':'Radishes', 'V':'Violets'}

class Garden:
    def __init__(self, diagram, students=full_student_list):
        self.diagram = diagram.split('\n')
        self.students = sorted(students)
    def plants(self, student):
        i = self.students.index(student)
        student_section = [*self.diagram[0][i*2:(i+1)*2], *self.diagram[1][i*2:(i+1)*2]]
        return [plants[plant] for plant in student_section]



