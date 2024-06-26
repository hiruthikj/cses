def weird_algorithm():
    n = int(input().strip())
    print_conjecture(n)
    print()


def print_conjecture(n):
    print(n, end=" ")
    if n == 1:
        return
    if n % 2 == 0:
        n //= 2
    else:
        n = n * 3 + 1
    print_conjecture(n)


if __name__ == "__main__":
    weird_algorithm()
