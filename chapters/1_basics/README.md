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
Example: func Hello and EuclidianDiv in [basics.go](basics.go)

- Functions are declared using the `func` keyword, followed by the function name.
- Arguments are defined using names first, types after. This is the opposite of the convention from the C family but similar to Python type hints.
- The return type is defined after the parameters (again, same as Python and opposite of C).
- Functions can return multiple results. Note: unlike in python, it is not a tuple.
- Functions can have named return values.

## Variables, constants & pointers
Example: func Swap and Perimeter in [basics.go](basics.go)

- Declare variable and constants with the `var` and `const` keyword respectively.
- The type comes after the name.
- A value can be assigned on declaration with `=`.
- If initialized, the type can be infered, so it can be omitted.
- Since it it common to initialize variables, an alternative syntax exists without `var` keyword or type: `:=`.
- Constants can only be basic types: character, string, boolean, or numeric values.
- An untyped numeric constant takes the type needed by its context.
- Variables and constants can be defined globally, though only with `var` or `const` syntax.
- Pointers are declared with a star before the type.
- There is no pointer arithmetic.
- The zero value of a pointer is `nil`.

## Conditions & loops
Example: func Min3, Countdown and EuclidianDiv in [basics.go](basics.go)

- Brackets are mandatory, parenthesis are not.
- The `for` keyword is used both for C-style "for" loops, "foreach" and "while" loops
- Like `for`, the `if` statement can start with a short statement to execute before the condition.
- Variables declared by short statements are restricted to the scope of the corresponding `if` or `for` block.

## Exercises
Code the functions in [exercise.go](exercise.go) so that the test written in [exercise_test.go](exercise_test.go) passes when invoked using:
```console
$ go test exercise.go exercise_test.go
```
