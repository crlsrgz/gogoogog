package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func readFileContent(filePath string) {

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}

func readFileByLine(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//read line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Printf("---> %s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func readFileByWord(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		fmt.Printf("-> %s\n", scanner.Text())
		count += 1
	}
	fmt.Println("word count is:", count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func readFileInChunks(filePath string, chunkSize int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	buf := make([]byte, chunkSize)

	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)

		}
		if err == io.EOF {
			break
		}
		fmt.Println(string(buf[:n]))
	}

}

func main() {
	readFileByWord("text.txt")
	readFileContent("text.txt")
	readFileByLine("text.txt")
	readFileInChunks("text.txt", 10)
}
