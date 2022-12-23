import random

n = 75

f = open(f'data_{n}.dat', 'w')
f.write(f'{n}\n')
for _ in range(n):
    f.write(f'{random.random()},{random.random()}\n')