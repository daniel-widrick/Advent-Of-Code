# WARNING!!! Uses try/except for lazy index bounds checking

filename = "input.txt"

with open(filename, 'r') as file:
    array =[list(line.strip()) for line in file]
    
    wordCount = 0
    for i in range(1,len(array)-1):
        for j in range(1, len(array[i])-1):
            if array[i][j] != 'A':
                continue;
            corners = ''.join([ array[i-1][j-1], array[i-1][j+1], array[i+1][j-1], array[i+1][j+1] ])
            if corners.count('M') == 2 and corners.count('S') == 2:
                if array[i-1][j-1] != array[i+1][j+1] and array[i+1][j-1] != array[i-1][j+1]: #ensure cross
                    wordCount+=1
    print(wordCount)

