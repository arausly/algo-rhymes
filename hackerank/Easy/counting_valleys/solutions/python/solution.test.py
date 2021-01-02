
from solution import counting_valleys


def test_counting_valleys():
    path1, steps1, paths2, step2 = "DDUUDDUDUUUD", 12, "UDDDUDUU", 8
    valleys1, valleys2 = counting_valleys(
        steps1, path1), counting_valleys(step2, paths2)
    assert(valleys1) == 2, "Expected the number of valleys to be 2"
    assert(valleys2) == 1, "Expected the number of valleys to be 1"


if __name__ == "__main__":
    test_counting_valleys()
    print("PASS \n ok")
