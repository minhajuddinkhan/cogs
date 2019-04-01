package cogs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

// TakeInput takes input from stdin
func TakeInput(msg string) string {

	blue := color.New(color.FgBlue)
	white := color.New(color.FgWhite)
	blue.Println(msg)
	reader := bufio.NewReader(os.Stdin)
	for {

		white.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		return text
	}
}

func GetUserAndPwd() (string, string, error) {

	username := TakeInput("Enter your username")
	fmt.Println("Enter Password")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", err
	}
	return username, string(bytePassword), nil
}
