package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func GetInput(filename string) ([]string, error) {

	path := fmt.Sprintf("./day1/%v.txt", filename)
	// 1. Open file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 2. Read line by line
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

var possibleVals = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

type NumInLine struct {
	Index int
	Data  string
}

func (n NumInLine) String() string {
	return fmt.Sprintf(`NumInLine{Index: %v,  Data: %v}`, n.Index, n.Data)
}

func Day1(filename string, reader func(string) ([]string, error)) int {
	lines, err := reader(filename)
	if err != nil {
		log.Fatal(err)
	}

	numbersInLines := make([][]NumInLine, 0)
	for _, line := range lines {
		occurences := make([]NumInLine, 0)
		for key, digit := range possibleVals {

			// I know it could be simpler. But found out the problem
			// previously is how I handle when a line contains
			// something like "oneeightwone"
			// by right, it should have 4 elements in 'occurences'
			// but due to issue not considering duplicate occurences
			// it only capture 3 elements because "one" only detects
			// index at first occurence
			if strings.Contains(line, key) {
				pattern := regexp.MustCompile(key)
				matches := pattern.FindAllStringIndex(line, -1)
				for _, match := range matches {
					index := match[0]
					data := NumInLine{
						Index: index,
						Data:  digit,
					}
					occurences = append(occurences, data)
				}
			}
		}
		sort.Slice(occurences, func(i, j int) bool {
			return occurences[i].Index < occurences[j].Index
		})
		// fmt.Println(occurences)
		numbersInLines = append(numbersInLines, occurences)
	}

	// 3. Get first occurence of first digit and last occurences of last digit
	calibrators := make([]string, 0)
	for _, line := range numbersInLines {
		if len(line) == 0 {
			continue
		} else if len(line) == 1 {
			// 3.1. If there is only 1 digit, then it's "<digit><digit>"
			calibrators = append(calibrators, line[0].Data+line[0].Data)
		} else {
			// 3.2. Concat first digit and last digit as "<first digit><last digit>"
			calibrators = append(calibrators, line[0].Data+line[len(line)-1].Data)
		}
	}

	// 4. Sum of all calibrators
	sum := 0
	for _, calibrator := range calibrators {
		v, _ := strconv.Atoi(calibrator)
		sum += v
	}

	return sum
}