class School:
    def __init__(self):
        self.grades = dict()
        self.students = []

    def add_student(self, name, grade):
        if name in self.roster():
            self.students.append(False)
        else:
            self.students.append(True)
            if not grade in self.grades:
                self.grades[grade] = []
            self.grades[grade].append(name)
            self.grades[grade].sort()

    def roster(self):
        result = []
        for grade in sorted(self.grades.keys()):
            result += self.grade(grade)
        return result

    def grade(self, grade_number):
        if grade_number in self.grades:
            return self.grades[grade_number]
        else: 
            return []

    def added(self):
        return self.students
