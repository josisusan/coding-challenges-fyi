package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	totalBytes := flag.Bool("c", false, "Count the number of bytes in a file")
	totalLines := flag.Bool("l", false, "Count the number of lines in a file")
	totalWords := flag.Bool("w", false, "Count the number of words in a file")
	totalChars := flag.Bool("m", false, "Count the number of characters in a file")

	flag.Parse()
	var filename string
	var data []byte
	var err error

	if len(flag.Args()) == 1 {
		filename = flag.Args()[0]
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening file (%s) \n", filename)
			return
		}
		defer file.Close()
		data, err = io.ReadAll(file)
		if err != nil {
			fmt.Printf("Error reading file")
			return
		}
	} else {
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Printf("Error reading Stdin")
		}
	}
	r := bytes.NewReader(data)
	if *totalBytes {
		byteCount, err := countTotalBytes(r)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		fmt.Printf("%d %s \n", byteCount, filename)
		return
	}
	if *totalLines {
		lineCount, err := countLines(r)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		fmt.Printf("%d %s \n", lineCount, filename)
		return
	}
	if *totalWords {
		wordCount, err := countWords(r)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		fmt.Printf("%d %s \n", wordCount, filename)
		return
	}
	if *totalChars {
		charCount, err := countCharacters(r)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		fmt.Printf("%d %s \n", charCount, filename)
	}

	if !(*totalBytes) && !(*totalLines) && !(*totalWords) && !(*totalChars) {
		lineCount, err := countLines(r)
		if err != nil {
			fmt.Printf("Error on line count: %v", err)
			return
		}
		r = bytes.NewReader(data)
		byteCount, err := countTotalBytes(r)
		if err != nil {
			fmt.Printf("Error on byte count: %v", err)
			return
		}
		r = bytes.NewReader(data)
		wordCount, err := countWords(r)
		if err != nil {
			fmt.Printf("Error on word count: %v", err)
			return
		}

		fmt.Printf("%d %d %d %s", lineCount, wordCount, byteCount, filename)
	}

}

func countTotalBytes(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanBytes)
	if count, err := totalCount(scanner); err != nil {
		return count, err
	} else {
		return count, nil
	}
}

func countLines(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	if count, err := totalCount(scanner); err != nil {
		return count, err
	} else {
		return count, nil
	}
}

func countWords(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	if count, err := totalCount(scanner); err != nil {
		return count, err
	} else {
		return count, nil
	}
}

func countCharacters(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)
	if count, err := totalCount(scanner); err != nil {
		return count, err
	} else {
		return count, nil
	}
}

func totalCount(scanner *bufio.Scanner) (int, error) {
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return count, err
	}
	return count, nil
}
