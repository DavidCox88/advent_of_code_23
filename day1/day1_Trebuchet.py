import re

def part_one(file):
    total = 0
    for x in file:
        t = re.sub("[^0-9]", "", x)
        n = t[0] + t[-1]
        n = int(n)
        total = total + n
    return print('Part one total is {}'.format(total))  

def part_two(file):
    dictionary = {"one": "1",
                  "two": "2",
                  "three": "3",
                  "four": "4",
                  "five": "5",
                  "six": "6",
                  "seven": "7",
                  "eight": "8",
                  "nine": "9"
                }
    f = open(r"puzzle_input.txt", "r")
    total = 0
    for x in f:

        for key in dictionary.keys():
            x = x.lower().replace(key, dictionary[key])
        t = re.sub("[^0-9]", "", x)
        n = t[0] + t[-1]
        n = int(n)
        total = total + n
    return print('Part two total is {}'.format(total))  

def main():
    file = open(r"puzzle_input.txt", "r")
    part_one(file)
    part_two(file)

if __name__ == "__main__":
    main()