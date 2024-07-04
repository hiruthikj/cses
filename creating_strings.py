from itertools import permutations

def creating_strings():
    inpStr = input().strip()

    charCount = {}
    for ch in inpStr:
        if ch in charCount:
            charCount[ch] += 1
        else:
            charCount[ch] = 1

    result = []
    permuteString(charCount, "", len(inpStr), result)
    result.sort()
    print(len(result))

    for s in result:
        print(s)

def permuteString(charCountRem, currStr, remCount, result):
    if remCount == 0:
        result.append(currStr)
        return

    for ch in list(charCountRem.keys()):
        charCountCp = charCountRem.copy()
        charCountCp[ch] -= 1
        if charCountCp[ch] == 0:
            del charCountCp[ch]
        permuteString(charCountCp, currStr + ch, remCount - 1, result)

# Example usage
creating_strings()

