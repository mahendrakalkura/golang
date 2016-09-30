from sys import argv

iterations = int(argv[1])

total = 0
for number in range(1, iterations + 1):
    total = total + number

print(total)
