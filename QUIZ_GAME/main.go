package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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
				continue
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

	reader := bufio.NewReader(os.Stdin)
	defer fd.Close()
	var score1 chan int
	csvReader := csv.NewReader(fd)

	duration := 5
	fmt.Print("Enter timer(Default is 5)   ::      ")
	duration1, _ := reader.ReadString('\n')
	duration1 = strings.TrimSpace(duration1)
	if duration1 != "" {
		duration, _ = strconv.Atoi(duration1)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(duration))

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
