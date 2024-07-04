def number_spiral():
    n = int(input())

    for _ in range(n):
        y, x = map(int, input().split())
        start = 0

        if y > x:
            if y % 2 == 0:
                start = (y * y) - (x - 1)
            else:
                start = (y - 1) * (y - 1) + x
        else:
            if x % 2 == 0:
                start = (x - 1) * (x - 1) + y
            else:
                start = x * x - y + 1
        
        print(start)

number_spiral()
