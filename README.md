# Pointers Project

This project demonstrates the use of pointers in Go. It includes examples of functions manipulating variables by their memory address, showcasing the effect of pointers on data modification.

## Features

- **Increment a Value**: Demonstrates incrementing an integer value using a pointer.
- **Swap Two Values**: Swaps the values of two integer variables using pointers.
- **Modify a Person's Age**: Modifies the age of a `Person` struct by using a pointer.

## Project Structure

- `main.go`: The entry point of the program that utilizes functions defined in the `pointers` package.
- `pkg/pointers.go`: Contains the definitions of functions that use pointers.
- `pkg/pointers_test.go`: Unit tests for the functions defined in `pointers.go`.

## Prerequisites

To run this project, you need to have Go installed on your machine. This project has been tested with Go version 1.22.5.

## How to Run

To run the main program:

```sh
go run main.go
```

## How to test each function

```sh
go test -v ./pkg -run TestIncrement
```

```sh
go test -v ./pkg -run TestSwap
```

```sh
go test -v ./pkg -run TestSetAge
```

## Example Usage

Here's a simple example demonstrating how to print a variable's value and its memory address:

```go
package main

import "fmt"

func main() {

  var num int = 5
  // prints the value stored in variable
  fmt.Println("Variable Value:", num)

  // prints the address of the variable
  fmt.Println("Memory Address:", &num)

}
```

Program to get the value pointed by a pointer

```go
package main
import "fmt"

func main() {

  var name = "John"
  var ptr *string

  ptr = &name

  // * to get the value pointed by ptr
  fmt.Println(*ptr) // John

}
```

Last example

```go
package main
import "fmt"

func main() {
  var num  int
  var ptr *int

  num = 22
  fmt.Println("Address of num:",&num)
  fmt.Println("Value of num:",num)

  ptr = &num
  fmt.Println("\nAddress of pointer ptr:",ptr)
  fmt.Println("Content of pointer ptr:",*ptr)

  num = 11
  fmt.Println("\nAddress of pointer ptr:",ptr)
  fmt.Println("Content of pointer ptr:",*ptr)

  *ptr = 2
  fmt.Println("\nAddress of num:",&num)
  fmt.Println("Value of num:",num)
}
```
