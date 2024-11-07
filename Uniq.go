package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	paramC := flag.Bool("c", false, "Количество встречаний строки во входных данных")
	paramD := flag.Bool("d", false, "Те строки, которые повторились во входных данных")
	paramU := flag.Bool("u", false, "Те строки, которые не повторились во входных данных.")
	paramI := flag.Bool("i", false, "Игнорировать регистр")
	paramF := flag.Int("f", 0, "Не учитывать первые num_fields полей в строке")
	paramS := flag.Int("s", 0, "Не учитывать первые num_chars символов в строке")

	flag.Parse()

	if *paramD && *paramU || *paramC && *paramU || *paramD && *paramC {
		fmt.Println("Параметры не могут использоваться одновременно!")
		return
	}

	inputFile := flag.Arg(0)
	outputFile := flag.Arg(1)

	var input *os.File
	var err error
	if inputFile == "" {
		input = os.Stdin
	} else {
		input, err = os.Open(inputFile)
		if err != nil {
			fmt.Println("Ошибка открытия файла!", err)
			return
		}
		defer input.Close()
	}

	var output *os.File
	if outputFile == "" {
		output = os.Stdout
	} else {
		output, err = os.Create(outputFile)
		if err != nil {
			fmt.Println("Ошибка создания файла!", err)
			return
		}
		defer output.Close()
	}

	scanner := bufio.NewScanner(input)
	inputLines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if *paramI {
			line = strings.ToLower(line)
		}

		if *paramF != 0 {
			fields := strings.Fields(line)
			if len(fields) > *paramF {
				line = strings.Join(fields[*paramF:], " ")
			} else {
				line = ""
			}
		}

		if *paramS != 0 && len(line) > *paramS {
			line = line[*paramS:]
		} else if len(line) <= *paramS {
			line = ""
		}

		if strings.TrimSpace(line) != "" {
			inputLines = append(inputLines, line)
		}
	}
	outputLines := []string{}
	if *paramC {
		countLine := 1
		prevLine := inputLines[0]
		for _, elem := range append(inputLines, " ")[1:] {
			if prevLine == elem {
				countLine++
			} else {
				outputLines = append(outputLines, strconv.Itoa(countLine)+" "+prevLine)
				countLine = 1
			}
			prevLine = elem
		}
	}
	if *paramD {
		countLine := 1
		prevLine := inputLines[0]
		for _, elem := range append(inputLines, " ")[1:] {
			if prevLine == elem {
				countLine++
			} else {
				if countLine > 1 {
					outputLines = append(outputLines, prevLine)
				}
				countLine = 1
			}
			prevLine = elem
		}
	}
	if *paramU {
		countLine := 1
		prevLine := inputLines[0]
		for _, elem := range append(inputLines, " ")[1:] {
			if prevLine == elem {
				countLine++
			} else {
				if countLine == 1 {
					outputLines = append(outputLines, prevLine)
				}
				countLine = 1
			}
			prevLine = elem
		}
	}
	for _, value := range outputLines {
		if strings.TrimSpace(value) != "" {
			if outputFile != "" {
				fmt.Fprintln(output, value)
			} else {
				fmt.Println(value)
			}
		}
	}
}
