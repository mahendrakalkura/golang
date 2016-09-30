from asyncio import get_event_loop
from concurrent.futures import as_completed, ThreadPoolExecutor
from pprint import pprint
from sys import argv

from requests import request

max_workers = int(argv[1])
iterations = int(argv[2])


def worker(number):
    try:
        response = request("GET", "https://bitbucket.org", timeout=5)
        return [response.status_code, len(response.content)]
    except Exception:
        pass
    return [999, 0]


async def main(max_workers, iterations):
    with ThreadPoolExecutor(max_workers=max_workers) as executor:
        futures = [
            executor.submit(worker, number)
            for number in range(1, iterations + 1)
        ]
        status_codes = {}
        size = 0
        for future in as_completed(futures):
            result = future.result()
            if result[0] not in status_codes:
                status_codes[result[0]] = 0
            status_codes[result[0]] += 1
            size = size + result[1]
        pprint(status_codes)
        print(size)

loop = get_event_loop()
loop.run_until_complete(main(max_workers, iterations))
loop.close()
