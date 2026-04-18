class Flower:
    def __init__(self, name, season):
        self.name = name
        self.season = season

    def describe(self):
        return f"{self.name} blooms in {self.season}."
