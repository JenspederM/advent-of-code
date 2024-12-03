from advent_of_code.utils import data_loader


def solve(example: str, part2=False):
    part_sum = 0
    i = 0
    mul_enabled = True
    mul = "mul("
    do = "do()"
    don_t = "don't()"
    l_mul = len(mul)
    l_do = len(do)
    l_don_t = len(don_t)
    while i < len(example):
        if mul_enabled and example[i : min(i + l_mul, len(example) - 1)] == mul:
            j = example.find(")", i)
            if (
                not all(c in "0123456789," for c in example[i + l_mul : j])
                or len(example[i + l_mul : j]) == 0
            ):
                i += 1
                continue
            try:
                x, y = example[i + l_mul : j].split(",")
                part_sum += int(x) * int(y)
                i = j + 1
            except Exception:
                i += 1
                raise Exception(
                    f"Error: {example[i + l_mul : j]}, {example[i-1 : j]} |||| {i}, {j}"
                )
        elif part2 and example[i : min(i + l_do, len(example) - 1)] == do:
            i = i + l_do
            mul_enabled = True
        elif part2 and example[i : min(i + l_don_t, len(example) - 1)] == don_t:
            i = i + l_don_t
            mul_enabled = False
        else:
            i += 1

    return part_sum


def main():
    lines = data_loader("day3.txt")
    print(sum(solve(line, part2=True) for line in lines))


if __name__ == "__main__":
    main()
