package handler

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
)

const dir = "taskFile"

func DateNow() string {
	out, _ := exec.Command("date", "+%F").Output()
	return string(bytes.TrimSpace(out))
}

func TimeNow() string {
	out, _ := exec.Command("date", "+%H:%M").Output()
	return string(bytes.TrimSpace(out))
}

func GetPath() string {
	return path.Join(dir, fmt.Sprintf("tasks_%s.json", DateNow()))
}

func Ask(text string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(text)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text())
	}
	return ""
}

func AskRequired(prompt string) (string, bool) {
	for {
		input := Ask(prompt)

		if input == "0" {
			return "", false
		}

		if input != "" {
			return input, true
		}
		fmt.Printf("[ERR] - Input cannot be empty. Try again.\n")
		ClearScreen()
		ListTasks()
	}
}

func AskIntRequired(prompt string) (int, bool) {
	for {
		input, ok := AskRequired(prompt)
		if !ok {
			return 0, false
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("[ERR] - Invalid number")
			ClearScreen()
			ListTasks()
			continue
		}
		return num, true
	}
}

func ValidateTime(input string) (string, error) {
	parts := strings.Split(input, ":")
	if len(parts) != 2 {
		return "", fmt.Errorf("Invalid format, use HH:MM")
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", fmt.Errorf("Invalid hour")
	}

	min, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", fmt.Errorf("Invalid minute")
	}

	if hour < 0 || hour > 23 {
		return "", fmt.Errorf("Hour must be 00 - 23")
	}

	if min < 0 || min > 59 {
		return "", fmt.Errorf("Minute must be 00 - 59")
	}

	return fmt.Sprintf("%02d:%02d", hour, min), nil
}

func ClearScreen() {
	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
}
