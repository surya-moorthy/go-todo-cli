package main

import (
	"bufio"
	"fmt"
	"go-mod-cli/storage"
	"go-mod-cli/todo"
	"os"
	"time"
)

// initiate todo service
// load json file and add it to todo service.
// when we close the application defer save json file.

var Storage storage.JsonStorage

type Todo struct {
	title       string
	description string
	createdAt   string
}

func main() {
	args := os.Args

	var storage storage.JsonStorage
	var todos []todo.Todo
	var service todo.Service

	if len(args) < 3 {
		fmt.Println("cmd must : ./main todo run")
		return
	}

	storage.Filepath = "data/data.json"
	storage.Load(&todos)
	service = todo.Service{}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Task Management cli")
		fmt.Println("Options : ")
		fmt.Print("1. add \n2. delete \n3. update \n4. display all todos \n5. search \n6. exit \nEnter the Option:")

		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error happened : ", err.Error())
			return
		}

		switch op := scanner.Text(); op {
		case "1":
			{

				fmt.Println("Adding todo")
				todo := todo.Todo{}

				fmt.Print("title : ")

				scanner.Scan()
				todo.Title = scanner.Text()

				fmt.Print("description : ")
				scanner.Scan()
				todo.Description = scanner.Text()

				todo.CreatedAt = time.Now()

				str, err := service.Add(todo)

				if err != nil {
					fmt.Printf(err.Error())
					break
				}

				fmt.Println(str)
			}

		case "2":
			{
				fmt.Println("Delete todo")

				fmt.Print("title : ")

				scanner.Scan()
				title := scanner.Text()

				str, err := service.Delete(title)

				if err != nil {
					fmt.Printf(err.Error())
					break
				}

				fmt.Println(str)
			}
		case "3":
			{
				fmt.Println("update")
			}
		case "4":
			{
				todos, err := service.List()
				if err != nil {
					fmt.Printf(err.Error())
					break
				}

				if len(todos) == 0 {
					fmt.Println("No todos are there.")
					break
				}

				fmt.Println("\n------------------------------")
				for _, todo := range todos {
					fmt.Printf("title : %s \ndescription : %s \n", todo.Title, todo.Description)
					fmt.Println("\n------------------------------")
				}
			}
		case "5": 
		{
			fmt.Println("search")
		}	
		case "6" : {
			os.Exit(1)
		}
		}

	}
}
