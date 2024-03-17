import copy
DriverData = {
    "HAM": {"price": 19.5, "Race1": 12, "Race2": 6},
    "NOR": {"price": 23.2, "Race1": 16, "Race2": 8},
    "PIA": {"price": 19.4, "Race1": 10, "Race2": 23},
    "ALO": {"price": 16.2, "Race1": 7, "Race2": 16},
    "STR": {"price": 11.0, "Race1": 8, "Race2": -17},
    "VER": {"price": 30.4, "Race1": 45, "Race2": 36},
    "ALB": {"price": 7.3, "Race1": 0, "Race2": 6},
    "SAR": {"price": 5.9, "Race1": 3, "Race2":7},
    "RIC": {"price": 8.7, "Race1": 5, "Race2": 0},
    "TSU": {"price": 7.6, "Race1": -1, "Race2": -1},
    "BOT": {"price": 6.0, "Race1": 0, "Race2": -1},
    "ZHO": {"price": 6.9, "Race1": 11, "Race2": -2},
    "HUL": {"price": 6.7, "Race1": -3, "Race2": 9},
    "MAG": {"price": 7.2, "Race1": 7, "Race2": 7},
    "OCO": {"price": 8.8, "Race1": 7, "Race2": 8},
    "GAS": {"price": 7.5, "Race1": 6, "Race2": -20},
    "PER": {"price": 22.1, "Race1": 31, "Race2": 31},
    "LEC": {"price": 20.4, "Race1": 22, "Race2": 37},
    "SAI": {"price": 18.8, "Race1": 36, "Race2": 0},
    "RUS": {"price": 19.2, "Race1": 20, "Race2": 15},
}

ConstructorData = {
    "Red Bull Racing": {"price": 28.2, "Race1": 89, "Race2": 90},
    "Ferrari": {"price": 19.9, "Race1": 73, "Race2": 58},
    "Mercedes": {"price": 20.3, "Race1": 42, "Race2": 36},
    "Mclaren": {"price": 23.6, "Race1": 36, "Race2": 41},
    "Aston Martin": {"price": 14, "Race1": 20, "Race2": 9},
    "Haas F1 Team": {"price": 6.7, "Race1": 9, "Race2": 19},
    "Williams": {"price": 6.7, "Race1": 4, "Race2": 14},
    "RB": {"price": 8.3, "Race1": 7, "Race2": 4},
    "Kick Sauber": {"price": 6.3, "Race1": 10, "Race2": -4},
    "Alpine": {"price": 8.1, "Race1": 12, "Race2": -13},
}

my_team = {
    "Driver1": None,
    "Driver2": None,
    "Driver3": None,
    "Driver4": None,
    "Driver5": None,
    "Constructor 1": None,
    "Constructor 2": None
}

def combinations(iterable, r=None):
    pool = tuple(iterable)
    n = len(pool)
    r = n if r is None else r
    if r > n:
        return
    indices = list(range(n))
    cycles = list(range(n, n-r, -1))
    yield tuple(pool[i] for i in indices[:r])
    while n:
        for i in reversed(range(r)):
            cycles[i] -= 1
            if cycles[i] == 0:
                indices[i:] = indices[i+1:] + indices[i:i+1]
                cycles[i] = n - i
            else:
                j = cycles[i]
                indices[i], indices[-j] = indices[-j], indices[i]
                yield tuple(pool[i] for i in indices[:r])
                break
        else:
            return

bestTeam = {"drivers": [],"constructors": []}
bestScore = 0

for combo in combinations(DriverData, 5):
    for combo2 in combinations(ConstructorData, 2):
        currentTeamPrice = 0
        currentTeamPoints = 0
        biggestDriver1 = 0
        biggestDriver2 = 0
        for driver in combo:
            if DriverData[driver]["Race1"] > biggestDriver1:
                biggestDriver1 = DriverData[driver]["Race1"]
            if DriverData[driver]["Race2"] > biggestDriver2:
                biggestDriver2 = DriverData[driver]["Race2"]
            currentTeamPrice += DriverData[driver]["price"]
            currentTeamPoints += DriverData[driver]["Race1"] + DriverData[driver]["Race2"]
        currentTeamPoints += biggestDriver1 + biggestDriver2
        for constructor in combo2:
            currentTeamPrice += ConstructorData[constructor]["price"]
            currentTeamPoints += ConstructorData[constructor]["Race1"] + ConstructorData[constructor]["Race2"]
        if currentTeamPoints > bestScore and currentTeamPrice <= 100:
            bestTeam['drivers'] = combo
            bestTeam['constructors'] = combo2
            bestTeamPrice = currentTeamPrice
            bestScore = currentTeamPoints
            print(bestTeam, bestTeamPrice, bestScore)




