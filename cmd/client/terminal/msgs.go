package terminal

import (
	"fmt"
	"gophertalk/internal/dto"
	"strings"
)

const clearCurrentLinePattern = "\r"

func printYourTime(to string) {
	if to == "" {
		to = "all"
	}
	fmt.Print(fmt.Sprintf("(type for %s):", to))
}

func printDefinedUser(to string) {
	if to == "" {
		to = "all"
	}
	fmt.Println(fmt.Sprintf("**Defined messages for %s", to))
}

func PrintHelp() {
	fmt.Println("------------------------HELP-----------------------")
	fmt.Println("/help\t\tto show the help message")
	fmt.Println("/users\t\tfor list connected users")
	fmt.Println("/to {user}\tto define the user to send the message.")
	fmt.Println("/all\t\tfor all people")
	fmt.Println("---------------------------------------------------")
}

func PrintNewReceivedMessage(obj dto.MessageDto) {
	if obj.Type == dto.MessageDtoTypeListUsers {
		printListOfUsers(obj)
	} else {
		printMessage(obj)
	}
}

func printListOfUsers(obj dto.MessageDto) {
	fmt.Println(clearCurrentLinePattern, "Users online:")
	fmt.Println(" ->", strings.Join(obj.Users, "\n -> "))
}

func printMessage(obj dto.MessageDto) {
	var to string
	if obj.To == "" {
		to = "for all"
	} else {
		to = "to you"
	}
	fmt.Println(clearCurrentLinePattern, obj.From, "say", to, ":", obj.Msg)
}
