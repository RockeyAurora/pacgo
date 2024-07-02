package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/danicat/simpleansi"
)

var maze []string

func loadMaze(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	for row, line := maze {
		for col, char := line {
			switch char  {
			case 'P' :
				player = sprite{row, col}
			}
		}
	}

	return nil
}

func printScreen() {
	simpleansi.ClearScreen()
	for _, line := range maze {
		fmt.Println(line)
	}
}

func readInput() (string, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	} else if cnt >= 3 {
		if buffer[0] ==0x1b && buffer[1] == '[' {
			switch buffer[2] {
			case 'A' :
				return "UP", nil
			case 'B' :
				return "DOWN", nil
			case 'C' :
				return "RIGHT", nil
			case 'D' :
				return "LEFT", nil
			}
		}
	}
	

	return "", nil
}

func makeMove(oldRow, oldCol int, dir string) (newRow, newCol int) {
	newRow, newCol = oldRow, oldCol
	switch dir {
	case "UP" :
		newRow = newRow - 1
		if newRow < 0 {
			newRow = len(maze) - 1
		}
	case "DOWN" :
		newRow = newRow + 1
		if newRow == len(maze) {
			newRow = 0
		}
	case "LEFT" :
		newCol = newCol - 1
		if newCol < 0 {
			newCol = len(maze) - 1
		}
	case "RIGHT" : {
		newCol = newCol + 1
		if newCol 
	}
	}

}

func initialise() {
	cbTerm := exec.Command("stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalln("unable to activate cbreak mode:", err)
	}
}

func cleanup() {
	cookedTerm := exec.Command("stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalln("unable to activate cooked mode:", err)
	}
}

type spwrite struct {
	row int 
	col int
}

var player sprite

func main() {
	// initialise game
	initialise()
	defer cleanup()

	// load resources
	err := loadMaze("maze01.txt")
	if err != nil {
		log.Println("failed to load maze:", err)
		return
	}

	// game loop
	for {
		// update screen
		printScreen()

		// process input
		input, err := readInput()
		if err != nil {
			log.Println("error reading input:", err)
			break
		}

		// process movement

		// process collisions

		// check game over
		if input == "ESC" {
			break
		}

		// repeat
	}
}