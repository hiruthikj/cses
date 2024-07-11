def trailing_zeroes():
    n = int(input().strip())

    factor_5_count = 0
    divisor = 5

    while divisor <= n:
        factor_5_count += n // divisor
        divisor *= 5

    print(factor_5_count)

# Example usage
trailing_zeroes()

