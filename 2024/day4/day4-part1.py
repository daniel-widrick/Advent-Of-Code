# WARNING!!! Uses try/except for lazy index bounds checking

filename = "input.txt"

with open(filename, 'r') as file:
    array =[list(line.strip()) for line in file]
    
    wordCount = 0
    for i in range(len(array)):
        for j in range(len(array[i])):
            if array[i][j] == 'X':
                # Search Right
                try:
                    right = ''.join([ array[i][j], array[i][j+1], array[i][j+2], array[i][j+3] ])
                    if right == 'XMAS':
                        print("R " + right)
                        wordCount=wordCount+1
                except:
                    pass
                try:
                    if j-3 < 0:
                        raise ValueError("Neative indexes are cancer!")
                    left = ''.join([ array[i][j], array[i][j-1], array[i][j-2], array[i][j-3] ])
                    if left == 'XMAS':
                        print("L " + left)
                        wordCount=wordCount+1
                except:
                    pass
                try:
                    down = ''.join([array[i][j], array[i+1][j], array[i+2][j], array[i+3][j]])
                    if down == 'XMAS':
                        print("D " + down)
                        wordCount=wordCount+1
                except:
                    pass #Ignore index exceptions
                try:
                    if i-3 < 0:
                        raise ValueError("Neative indexes are cancer!")
                    up = ''.join([array[i][j], array[i-1][j], array[i-2][j], array[i-3][j]])
                    if up == 'XMAS':
                        print("U " + up)
                        wordCount=wordCount+1
                except:
                    pass #Ignore index exceptions

                ### Handle Diagonals ##DIRTY But simple
                try:
                    if i-3 < 0 or j-3 < 0:
                        raise ValueError("Neative indexes are cancer!")
                    upLeft = ''.join([ array[i][j], array[i-1][j-1], array[i-2][j-2], array[i-3][j-3] ])
                    if upLeft == 'XMAS':
                        print("UL " + upLeft)
                        wordCount=wordCount+1
                except:
                    pass #still ignore index exceptions
                try:
                    if i-3 < 0:
                        raise ValueError("Neative indexes are cancer!")
                    upRight = ''.join([ array[i][j], array[i-1][j+1], array[i-2][j+2], array[i-3][j+3] ])
                    if upRight == 'XMAS':
                        print("UR " + upRight)
                        wordCount=wordCount+1
                except:
                    pass
                try:
                    if j-3 < 0:
                        raise ValueError("Neative indexes are cancer!")
                    downLeft = ''.join([ array[i][j], array[i+1][j-1], array[i+2][j-2], array[i+3][j-3] ])
                    if downLeft == 'XMAS':
                        print("DL " + downLeft)
                        wordCount=wordCount+1
                except:
                    pass
                try:
                    downRight = ''.join([ array[i][j], array[i+1][j+1], array[i+2][j+2], array[i+3][j+3] ])
                    if downRight == 'XMAS':
                        print("DR " + downRight)
                        wordCount=wordCount+1
                except:
                    pass



    print(wordCount)

