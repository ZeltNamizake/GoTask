package handler

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const dir = "taskFile"

func DateNow() string {
	out, _ := exec.Command("date", "+%F").Output()
	return string(bytes.TrimSpace(out))
}

func GetPath() string {
	return fmt.Sprintf("%s/tasks_%s.json", dir, DateNow())
}

func Ask(text string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(text)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text())
	}
	return ""
}

func AskRequired(prompt string) string {
	for {
		input := Ask(prompt)
		if input != "" {
			return input
		}
		fmt.Printf("[ERR] - Input cannot be empty. Try again.\n")
		ClearScreen()
	}
}

func AskIntRequired(prompt string) int {
	for {
		input := AskRequired(prompt)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("[ERR] - Invalid number")
			ClearScreen()
			continue
		}
		return num
	}
}

func ClearScreen() {
	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
}
