package main

import (
	"bufio"
	"fmt"
	"go-mod-cli/model"
	"go-mod-cli/storage"
	"go-mod-cli/todo"
	"os"
	"strings"
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

	var Storage storage.JsonStorage
	var todos []model.Todo
	var Service todo.Service

	if len(args) < 3 {
		fmt.Println("cmd must : ./main todo run")
		return
	}

	Storage.Filepath = "./data/data.json"
	todos, err := Storage.Load()

	if err != nil {
		 fmt.Println(err.Error())
		 return
	}

	Service = todo.Service{}

	Service.Storage = &Storage
	Service.Todos = todos

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
				todo := model.Todo{}

				fmt.Print("title : ")

				scanner.Scan()
				todo.Title = scanner.Text()

				fmt.Print("description : ")
				scanner.Scan()
				todo.Description = scanner.Text()

				todo.CreatedAt = time.Now()

				str, err := Service.Add(todo)

				if err != nil {
					fmt.Printf("%s", err.Error())
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

				str, err := Service.Delete(title)

				if err != nil {
					fmt.Printf("%s", err.Error())
					break
				}

				fmt.Println(str)
			}
		case "3":
			{
				fmt.Print("todo Title to update : ")
				scanner.Scan()
				title := scanner.Text()
				todo, err := Service.Search(title)

				if err != nil {
					fmt.Printf("%s", err.Error())
					break
				}

				fmt.Println("which one do you wanna update : ")
				fmt.Print("1. title only \n2. description only \n3.both \nEnter the Option:")

				scanner.Scan()
				switch up := scanner.Text(); up { 
				case "1":
					fmt.Print("Enter the new title : ")
					scanner.Scan()
					todo.Title = strings.TrimSpace(scanner.Text())

				case "2":
					fmt.Print("Enter the new description : ")
					scanner.Scan()
					todo.Description = strings.TrimSpace(scanner.Text())
				
				case "3":
					fmt.Print("Enter the new title : ")
					scanner.Scan()
					todo.Title = strings.TrimSpace(scanner.Text())
					
					fmt.Print("Enter the new description : ")
					scanner.Scan()
					todo.Description = strings.TrimSpace(scanner.Text())
				default :
					fmt.Println("No such option.")
				}

				str, updateErr := Service.Update(title, todo)

				if updateErr != nil {
					fmt.Printf("%s", updateErr.Error())
					break
				}

				fmt.Println(str)
			}
		case "4":
			{
				todos, err := Service.List()
				if err != nil {
					fmt.Printf("%s", err.Error())
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
			fmt.Print("todo Title to update : ")
			scanner.Scan()
			title := scanner.Text()
			todo, err := Service.Search(title)

			if err != nil {
				fmt.Printf("%s", err.Error())
				break
			}

			fmt.Println("\n------------------------------")
			fmt.Printf("title : %s \ndescription : %s \n", todo.Title, todo.Description)
			fmt.Println("\n------------------------------")
		}	
		case "6" : {
			os.Exit(1)
		}
		}

	}
}
