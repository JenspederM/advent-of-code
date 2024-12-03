def test_part1():
    test = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
    expected = 161

    from advent_of_code.day3 import solve

    assert solve(test) == expected


def test_part2():
    test = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
    expected = 48

    from advent_of_code.day3 import solve

    assert solve(test, part2=True) == expected


def test_solve_open_ended():
    test = (
        "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))mul("
    )
    expected = 48

    from advent_of_code.day3 import solve

    assert solve(test, part2=True) == expected


def test_solve_open_ended_mid():
    test = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(mul(32,64](mul(11,8)undo()?mul(8,5))mul("
    expected = 48

    from advent_of_code.day3 import solve

    assert solve(test, part2=True) == expected


def test_part2_open_ended_mid():
    test = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(mul()mul(32,64](mul(11,8)undo()?mul(8,5))mul("
    expected = 48

    from advent_of_code.day3 import solve

    assert solve(test, part2=True) == expected
