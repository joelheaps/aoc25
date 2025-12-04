from dataclasses import dataclass
import math
from pathlib import Path
from collections import defaultdict
import sys


@dataclass(frozen=True)
class Junction:
    x: float
    y: float
    z: float

    def dist_from(self, other: "Junction") -> float:
        return math.sqrt(
            math.pow(self.x - other.x, 2)
            + math.pow(self.y - other.y, 2)
            + math.pow(self.z - other.z, 2)
        )


def get_closest_pairs(junctions: list[Junction]) -> list[tuple[Junction, Junction]]:
    pairs: dict[frozenset[Junction], float] = defaultdict(lambda: float("inf"))
    # Get closest neighbor to each junction
    for ji in junctions:
        for jj in junctions:
            if ji == jj:
                continue
            pairs[frozenset((ji, jj))] = ji.dist_from(jj)

    # Sort
    sorted_pairs = sorted(pairs.items(), key=lambda item: item[1])

    return [tuple(p[0]) for p in sorted_pairs]


def main() -> None:
    if sys.argv[1] == "part1":
        assert len(sys.argv) == 4, "specify part, input file, and connection count"
        part1()
    elif sys.argv[1] == "part2":
        part2()
    else:
        print("invalid part")


"""
I used a double dictionary method to keep track of Junctions and Circuits -- one tracks
which Junctions are members of a Circuit, and the other tracks which Circuit a particular
Junction belongs to.  They are updated together.
"""

def part1() -> None:
    input = Path(sys.argv[2]).read_text()
    count = int(sys.argv[3])
    raw_junctions = input.splitlines()
    junctions = []
    for jr in raw_junctions:
        if not jr:
            continue
        xyz = jr.split(",")
        junctions.append(Junction(int(xyz[0]), int(xyz[1]), int(xyz[2])))
    clostest_pairs = get_closest_pairs(junctions)

    circuits: dict[str, set[Junction]] = dict()
    junction_membership: dict[Junction, str] = dict()

    # Initialize circuits
    for i, j in enumerate(junctions):
        circuits[i] = {j}
        junction_membership[j] = i

    # Add junctions to circuits, in order
    for pair in clostest_pairs[:count]:
        left = pair[0]
        right = pair[1]

        # Combine circuits (use left's circuit as target).
        # All junctions are in a circuit
        target = junction_membership[left]
        other = junction_membership[right]
        if target == other:
            continue
        other_members = circuits.pop(other)
        circuits[target] = circuits[target] | other_members

        # Update circuit membership
        for om in other_members:
            junction_membership[om] = target

    # Process circuit size and print summary
    sizes: dict[int, int] = defaultdict(lambda: 0)
    for c in circuits.values():
        sizes[len(c)] += 1

    top3 = sorted(sizes.keys())[::-1][:3]
    print(math.prod(top3))


def part2() -> None:
    input = Path(sys.argv[2]).read_text()
    raw_junctions = input.splitlines()
    junctions = []
    for jr in raw_junctions:
        if not jr:
            continue
        xyz = jr.split(",")
        junctions.append(Junction(int(xyz[0]), int(xyz[1]), int(xyz[2])))
    clostest_pairs = get_closest_pairs(junctions)

    circuits: dict[str, set[Junction]] = dict()
    junction_membership: dict[Junction, str] = dict()

    # Initialize circuits
    for i, j in enumerate(junctions):
        circuits[i] = {j}
        junction_membership[j] = i

    # Add junctions to circuits, in order
    for pair in clostest_pairs:
        left = pair[0]
        right = pair[1]

        # Combine circuits (use left's circuit as target).
        # All junctions are in a circuit
        target = junction_membership[left]
        other = junction_membership[right]
        if target == other:
            continue
        other_members = circuits.pop(other)
        if len(circuits) == 1:
            print("Last connection: ", pair)
            print("X coord product: ", left.x * right.x)
            return
        circuits[target] = circuits[target] | other_members

        # Update circuit membership
        for om in other_members:
            junction_membership[om] = target

    # Process circuit size and print summary
    sizes: dict[int, int] = defaultdict(lambda: 0)
    for c in circuits.values():
        sizes[len(c)] += 1

    top3 = sorted(sizes.keys())[::-1][:3]
    print(math.prod(top3))


if __name__ == "__main__":
    main()
