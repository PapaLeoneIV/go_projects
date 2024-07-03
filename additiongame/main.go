package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"errors"
	"bufio"
)

func printresults(r int, w int, n int){
	fmt.Println("\n****DRUM ROLL****")
	fmt.Println("\nTotal number of questions: ", n)
	fmt.Println("\nNumber of correct answers: ", r)
	fmt.Println("\nNumber of wrong answers: ", w)
	fmt.Println("\nTHANKS FOR PLAYING!!!")
}


func readfile(file *os.File) (string, error) {
	var filestream []byte
	buf := make([]byte, 100)

	for {
		n, err := file.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return "", fmt.Errorf("error reading file: %v", err)
		}
		filestream = append(filestream, buf[:n]...)
	}
	return string(filestream), nil
}

func promptuser(key string, n int, value string){ 
	fmt.Println("Question n° "+  strconv.Itoa(n) + ":\nSolve this -->", key)
}

func getuserprompt(chan ch1) (int, error) {

	var buf string

	/* start:= time.Now() */
	go fmt.Scanln(&buf)
	go for {
		start := <-ch1
	if time.Now() - start > 6 * time.Second{
		fmt.Println("You took to long to answear!")
		os.Exit()
		}
	}
	out , err :=strconv.Atoi(buf)
	if err != nil{
		return 0, errors.New("User input: not a valid input!")
	} 
	return out, nil
}

func parse_csvfile(file *os.File) map[string]string{

	mapp := make(map[string]string)
	scanner := bufio.NewScanner(file) 

	for scanner.Scan() {
		line := scanner.Text()
		keyval := strings.Split(line, ",")
		if len(keyval) != 2 {
			fmt.Println("Wrong number of arguments in the equation!")
			continue
		}
		mapp[keyval[0]] = keyval[1]
	}
	return mapp
}

func handle_game_logic(mapp map[string]string, chan ch1) (right int, wrong int, n int){
	n = 0

	for k, val := range mapp {

		ch1 <- time.Now()
		promptuser(k, n, val)

		ans, err := getuserprompt(ch1)
		time := <-ch1
		if err != nil {
			fmt.Println("Invalid input. Next question!")
			wrong++
		} else {
			corr_ans, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("Invalid problem answer format in file.")
				continue
			}
			if ans == corr_ans {
				right++
			} else {
				wrong++
			}
		}
		n++
	}

	return right, wrong, n
}



func main() {

	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		return
	}
	defer file.Close()
	
	mapp := parse_csvfile(file)

	osR :=	bufio.NewReader(os.Stdin)
	fmt.Println("Press ENTER to start the quiz...")
	osR.ReadString('\n')

	ch1 := make([]chan , 2)
	
	right, wrong, n_quest := handle_game_logic(mapp, ch1)
	printresults(right, wrong, n_quest)
}
