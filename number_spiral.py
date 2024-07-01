def number_spiral():
    n = int(input())
    
    for _ in range(n):
        y, x = map(int, input().split())
        if y > x:
            start = (y % 2) + (y - (y % 2)) ** 2
            if y % 2 == 0:
                start -= x - 1
            else:
                start += x - 1
        else:
            start = (x - 1 + (x % 2)) ** 2 + 1 - (x%2)
            if x % 2 == 0:
                start += y - 1
            else:
                start -= y - 1
        
        print(start)

number_spiral()
