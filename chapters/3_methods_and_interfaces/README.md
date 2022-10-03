# Methods and interfaces

## Packages & imports

- Every Go program is made up of packages.
- Programs start running in package main. 
- A name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.
- pizza and pi do not start with a capital letter, so they are not exported.
- When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package. 
```go
// Exported
const Pi = 3.14
const Pizza = "marguarita"
// Not exported
const pi = 3.14
const pizza = "marguarita"
```

## Methods

- Go does not have classes. However, you can define methods on types.
- A method is a function with a special receiver argument.
- The receiver appears in its own argument list between the func keyword and the method name. 
- You can only declare a method with a receiver whose type is defined in the same package as the method.
- You cannot declare a method with a receiver of a built-in type such as int.
- If you want to "extend" a built-in type or a type from another package, you have to define your own type as an alias.
- You can declare methods with pointer receivers.
```go
type MyInt int
func (i MyInt) Successor() int {
    return i + 1
}

type Rect struct{
    height, width float64
}
func (r *Rect) Area() float64 {
    return r.height * r.width
}

func main() {
    fmt.Println(MyInt(3).Successor())
    r := &Rect{3, 4}
    fmt.Println(r.Area())
}
```

## Interface

- An interface type is defined as a set of method signatures.
- A value of interface type can hold any value that implements those methods.
- A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword. 
- The interface type that specifies zero methods is known as the empty interface: `interface{}`. An empty interface may hold values of any type.
- The built-in type `any` is an alias for the empty interface.
- `comparable` is a built-in interface that makes it possible to use the `==` and `!=` operators on values of the type.
- A type assertion `i.(T)` provides access to an interface value's underlying concrete value.
```go
var i interface{} = "hello"

// Asserts that the interface value i holds the concrete type string
// and assigns the underlying string value to the variable s.
s := i.(string) // "hello"
fmt.Println(s)

s, ok := i.(string) // "hello", true
fmt.Println(s, ok)

f, ok := i.(float64) // 0, false
fmt.Println(f, ok)

f = i.(float64) // panic
fmt.Println(f)
```

## Exercise: Mortals and DeathListeners

Create the Notary and Aristocrat structs and define their methods so that the tests passes.
They will have to implement the interfaces from [death_listener.go](death_listener.go).
The Notary listens for a Aristocrat's death and proceeds to share its legacy among its heirs.
An Aristocrats that fell in love listens to its soulmate's death and kills itself, Romeo and Juliet style.
