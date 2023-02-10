package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	fd, error := os.Open("problems.csv")
	// scanner := bufio.NewReader(os.Stdin)
	var ans int
	score := 0
	if error != nil {
		log.Fatal("Error in opening the file")
	}
	defer fd.Close()

	csvReader := csv.NewReader(fd)

	for {
		line, error := csvReader.Read()
		if error == io.EOF {
			break
		}
		if error != nil {
			log.Fatal("File error")
		} else {
			fmt.Print(line[0])
			fmt.Print("=  ")
			fmt.Scanf("%d\n", &ans)
			g := line[1]
			g1, _ := strconv.Atoi(g)
			if ans != g1 {
				break
			} else {
				score++
			}
		}

	}
	fmt.Println(score)
}
