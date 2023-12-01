package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
    // open file
	f, err := os.Open("puzzle_input.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file line by line using scanner
    scanner := bufio.NewScanner(f)
	total := 0
    for scanner.Scan() {
        // replace everything that isn't a number
		reg := regexp.MustCompile("[^0-9]")
		t := scanner.Text()
		t = reg.ReplaceAllString(t, "")
		n := t[0:1]
		n2  := t[len(t)-1:]
		n3 := n + n2
		s, err := strconv.Atoi(n3);
		if err != nil {
			fmt.Println("Can't convert this to an int!")
		} else {
			total = total + s
		}
    }
	fmt.Println("Total for part one is", total)

	f, err = os.Open("puzzle_input.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file line by line using scanner
    scanner2 := bufio.NewScanner(f)
	total2 := 0
    for scanner2.Scan() {
        // replace text first, then replace everything that isn't a number
		t := scanner2.Text()
		t = strings.Replace(t, "one", "o1e", -1)
		t = strings.Replace(t, "two", "t2o", -1)
		t = strings.Replace(t, "three", "t3e", -1)
		t = strings.Replace(t, "four", "f4r", -1)
		t = strings.Replace(t, "five", "f5e", -1)
		t = strings.Replace(t, "six", "s6x", -1)
		t = strings.Replace(t, "seven", "s7n", -1)
		t = strings.Replace(t, "eight", "e8t", -1)
		t = strings.Replace(t, "nine", "n9e", -1)
		reg := regexp.MustCompile("[^0-9]")
		t = reg.ReplaceAllString(t, "")
		n := t[0:1]
		n2  := t[len(t)-1:]
		n3 := n + n2
		s, err := strconv.Atoi(n3);
		if err != nil {
			fmt.Println("Can't convert this to an int!")
		} else {
			total2 = total2 + s
		}
    }
	fmt.Println("Total for part two is", total2)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}