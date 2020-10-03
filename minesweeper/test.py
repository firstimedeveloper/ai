tup = set(tuple((i,j) for i in range(0,8) for j in range(0,8)))
mines = set()
mines.add((4,0))
new = tup - mines
print(len(tup))
print(len(new))
print(next(iter(tup-mines)))