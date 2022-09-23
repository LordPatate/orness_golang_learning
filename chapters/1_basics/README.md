# Basics

## How to "hello world"
Example: [hello.go](hello.go)

- Write in a `*.go` file. We'll use `hello.go` for this example.
- Declare it as `package main` to be able to run it as a program.
- Import the _fmt_ package from the standard library to have access to printing functions.
- Use a `main` function as entry point.
- Print a single line with the [fmt.Println](https://pkg.go.dev/fmt#Println) function.
- The [go run](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program) command allows to both compile and run at once: `go run hello.go`.
You can also compile first using `go build hello.go` then run the generated executable (`hello.exe` on windows, `hello` on linux...).

## Functions
- Functions are declared using the `func` keyword, followed by the function name.
- Arguments are defined using names first, types after. This is the opposite of the convention from the C family but similar to Python type hints.
- The return type is defined after the parameters (again, same as Python and opposite of C).
- Functions can return multiple results. Note: unlike in python, it is not a tuple.
## 