from concurrent.futures import as_completed, ThreadPoolExecutor
from sys import argv


def worker(number):
    return number

max_workers = int(argv[1])
iterations = int(argv[2])

with ThreadPoolExecutor(max_workers=max_workers) as executor:
    futures = [
        executor.submit(worker, number) for number in range(1, iterations + 1)
    ]
    total = 0
    for future in as_completed(futures):
        total = total + future.result()
    print(total)
