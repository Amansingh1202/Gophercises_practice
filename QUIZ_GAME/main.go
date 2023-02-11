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
	var ans string
	reader := bufio.NewReader(os.Stdin)
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
			ans, _ = reader.ReadString('\n')
			ans = strings.TrimSpace(ans)
			ans = strings.ToLower(ans)
			g := strings.ToLower(line[1])
			if ans != g {
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
