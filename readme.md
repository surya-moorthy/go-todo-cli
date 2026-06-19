# Todo Management CLI

- An CLI tool that has basic crud operations on todos.
- Each todo has a title and description.
- this tool is built not because it is easier but to build a CLI in golang with necessary fundamentals.

## How to Run?
- give a command ```go run main.go todo run``` in the terminal. 

## Operations
- **Add** : adds a todo by getting title and description as input data.
- **Update** : updates a todo by getting title of the todo the user wants to update and new data.
- **Delete** : deletes a todo with matching given title.
- **List** : list all the todos.
- **Search** : search todo with the given title.

## Versions

In this repo we have versions that are differentiated has branches, the **main** branch has the recent version.

### V1
- v1 is a basic brute approach. i used:
    1. os package for std cli.
    2. for loop, structs, methods
    3. in memory array, that stores todos temporarily until the program runs.

### V2
- v2 is a much focused on architecture and persistant storage.
    1. Package organization.
    2. Handling json file for persistant storage.
