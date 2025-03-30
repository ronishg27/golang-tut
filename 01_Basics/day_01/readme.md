# GoLang Basics - Day 01

## Basic setup for GoLang

- create a new folder
- run a command `go mod init [username/project_name]` to initialize a new module
- create a new file with the name `main.go`
- add the following code to the file

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

- in the terminal on the folder where the file is located, run the following command `go run main.go`

## Print Types

In go there are three ways to print a value to the console

1. `fmt.Println` - prints a value and a new line
2. `fmt.Printf` - prints a value in the format specified
3. `fmt.Print` - prints a value without a new line
