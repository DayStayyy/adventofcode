import re

f = open("./day5/data.txt", "r")
data = f.read()
data = data.split("\n")


crate = [
    ["B","Z","T"],
    ["V","H","T","D","N"],
    ["B","F","M","D"],
    ["T","J","G","W","V","Q","L"],
    ["W","D","G","P","V","F","Q","M"],
    ["V","Z","Q","G","H","F","S"],
    ["Z","S","N","R","L","T","C","W"],
    ["Z","H","W","D","J","N","R","M"],
    ["M","Q","L","F","D","S"]
]

crate_part_two = [
    ["B","Z","T"],
    ["V","H","T","D","N"],
    ["B","F","M","D"],
    ["T","J","G","W","V","Q","L"],
    ["W","D","G","P","V","F","Q","M"],
    ["V","Z","Q","G","H","F","S"],
    ["Z","S","N","R","L","T","C","W"],
    ["Z","H","W","D","J","N","R","M"],
    ["M","Q","L","F","D","S"]
]



for instruction in data :
    instruction_array = re.split('move|from |to ', instruction)
    instruction_array.pop(0)
    instruction_array = [ int(x)-1 for x in instruction_array]
    instruction_array[0] += 1
    for i in range(instruction_array[0]) :
        if(len(crate[instruction_array[1]]) > 0) :
            crate[instruction_array[2]].append(crate[instruction_array[1]][-1])
            crate[instruction_array[1]].pop()

    if(len(crate_part_two[instruction_array[1]]) > 0):
        if(len(crate_part_two[instruction_array[1]]) > instruction_array[0]) :
            crate_part_two[instruction_array[2]].extend(crate_part_two[instruction_array[1]][len(crate_part_two[instruction_array[1]])-instruction_array[0]:])
            crate_part_two[instruction_array[1]] = crate_part_two[instruction_array[1]][0:len(crate_part_two[instruction_array[1]])-instruction_array[0]]
        else :
            crate_part_two[instruction_array[2]].extend(crate_part_two[instruction_array[1]][0:])
            crate_part_two[instruction_array[1]] = []





result = ""
for column in crate: 
    if(len(column) != 0):
        result += column[-1]

result2 = ""
for column in crate_part_two: 
    if(len(column) != 0):
        result2 += column[-1]

print("Part 1 : \nThe top crate is ",result)
print("Part 1 : \nnThe top crate is ",result2)