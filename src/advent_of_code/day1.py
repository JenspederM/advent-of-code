from advent_of_code.utils import data_loader
import re


def main():
    lines = data_loader("day1.txt")
    pairs = [[int(v) for v in re.split(" +", line)] for line in lines]
    columns = [sorted(location_id) for location_id in zip(*pairs)]
    print(sum([max(a, b) - min(a, b) for a, b in zip(*columns)]))


if __name__ == "__main__":
    main()
