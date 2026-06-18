package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Todo struct {
	title       string
	description string
	createdAt   string
}

var todos []Todo

func (t *Todo) Add() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Adding todo")

	fmt.Print("title : ")
	scanner.Scan()
	t.title = scanner.Text()

	fmt.Print("description : ")
	scanner.Scan()
	t.description = scanner.Text()

	t.createdAt = time.Now().String()

	todos = append(todos, *t)

	return "successfully added.", nil
}

func Search(input string) (n int, ok bool) {
	n = 0
	ok = false

	for i, todo := range todos {
		if input == todo.title {
			n , ok = i, true
			break 
		}
	}

	return
}

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("cmd must : ./main todo run")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("Task Management cli")
		fmt.Println("Options : ")
		fmt.Print("1. add \n2. delete \n3. update \n4. display all todos \n5. exit \nEnter the Option:")

		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error happened : ", err.Error())
			return
		}

		switch op := scanner.Text(); op {

		case "1":

			var todo Todo
			str, err := todo.Add()

			if err != nil {
				fmt.Println(err.Error())
			}

			fmt.Println(str)

		case "2":
			fmt.Print("Enter the title of the todo to delete :")
			
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())
			idx, ok := Search(title)

			if !ok {
				fmt.Println("todo not found.")
				break
			}

			todos = append(todos[:idx], todos[idx + 1:]...)
			
			fmt.Println("\nsuccessfully deleted.")
		case "3":
			fmt.Print("Enter the title of the todo to update :")
			
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())
			idx, ok := Search(title)

			if !ok {
				fmt.Println("todo not found.")
				break
			}
			for {

				fmt.Println("which one do you wanna update : ")
				fmt.Print("1. title only \n2. description only \n3.both \nEnter the Option:")

				ok := true

				scanner.Scan()
				switch up := scanner.Text(); up { 
				case "1":
					fmt.Print("Enter the new title : ")
					scanner.Scan()
					todos[idx].title = strings.TrimSpace(scanner.Text())

				case "2":
					fmt.Print("Enter the new description : ")
					scanner.Scan()
					todos[idx].description = strings.TrimSpace(scanner.Text())
				
				case "3":
					fmt.Print("Enter the new title : ")
					scanner.Scan()
					todos[idx].title = strings.TrimSpace(scanner.Text())
					
					fmt.Print("Enter the new title : ")
					scanner.Scan()
					todos[idx].description = strings.TrimSpace(scanner.Text())

				default:
					ok = false
					fmt.Println("no such option.")
				}

				if ok {
					fmt.Println("\nsuccessfully updated.")
					break
				}
			
			}

		case "4":
			fmt.Println("\n------------------------------")
			for _, todo := range todos {
				fmt.Printf("title : %s \ndescription : %s \n", todo.title, todo.description)
				fmt.Println("\n------------------------------")
			}
		case "5":
			os.Exit(1)
		default:
			fmt.Println("no such option.")
		}
	}

}
