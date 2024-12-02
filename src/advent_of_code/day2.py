from advent_of_code.utils import data_loader


def is_valid(report):
    pairs = list(zip(report, report[1:]))
    diffs = [b - a for a, b in pairs]
    if diffs[0] < 0:
        _valid = all(diff < 0 and diff >= -3 and diff <= -1 for diff in diffs)
    else:
        _valid = all(diff > 0 and diff >= 1 and diff <= 3 for diff in diffs)
    return _valid


def main():
    lines = data_loader("day2.txt")
    reports = [[int(v) for v in line.split(" ")] for line in lines]
    # print("\n".join(str(report) for report in reports))
    print("result", sum(is_valid(report) for report in reports))


if __name__ == "__main__":
    main()
