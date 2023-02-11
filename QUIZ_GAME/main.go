package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var score = 0

func QuizGame(csvReader *csv.Reader, score1 chan int) {
	var ans int
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
	score1 <- score
}
func main() {
	fd, error := os.Open("problems.csv")
	if error != nil {
		log.Fatal("Error in opening the file")
	}
	defer fd.Close()
	var score1 chan int
	csvReader := csv.NewReader(fd)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	go QuizGame(csvReader, score1)

	select {
	case v := <-score1:
		fmt.Println(v)
	case <-ctx.Done():
		fmt.Println()
		fmt.Println("Time's up")
		fmt.Println(score)
		return
	}

}
