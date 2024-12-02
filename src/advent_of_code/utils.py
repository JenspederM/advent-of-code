from __future__ import annotations
from importlib.resources import files

package_files = files("advent_of_code.data")


def data_loader(path: str) -> list[str]:
    with package_files.joinpath(path).open() as f:
        return f.read().splitlines()
