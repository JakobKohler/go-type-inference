# Type inference

### What is Type Inference in Go?

Type inference in Go is the ability of the Go compiler to automatically determine the type of a variable based on the value assigned to it. This means that programmers do not have to explicitly state the type of a variable. For example, when a variable is declared using the short variable declaration syntax (e.g., `x := 10`), the compiler infers that `x` is of type `int` because the assigned value is an integer. This feature helps make the code cleaner and easier to read.

### Why is Type Inference Needed?

First, it makes coding faster by reducing the amount of code that developers need to write. This can help prevent mistakes that might happen when specifying types manually. Second, it improves code readability, making it easier for developers to understand what the code is doing. By letting the compiler figure out the types, Go allows for a more straightforward coding style. Finally, type inference keeps the strong typing of Go, ensuring that type safety is maintained while still making programming more convenient.

### How Type Inference works in Go
The general process of type inference in Go is to recursively compare the types of two expressions with each other until they are identical in a process called unification and then expand the results for the remaining type parameters. The following is going to explore what this means precisely and showcase how its done in Go

#### 0. Notiation

#### 1. Unification
In computer science the term unification is generally used to describe an algorithmic process of solving equations between expressions by replacing their variables with suitable terms until the resulting expressions are equivalent.

In the context of type inference, these equations are type equations derived from the type arguments or the assignments. There are several ways of collecting these type equations from the code to allow of the types to be inferred. The most trivial source of a type equation is the equations given by explicitly provided type arguments like `foo[List] (...)` for the generic function `func foo[A any]()...`. This results in the trivial equation `A ≡ List`. Another way to obtain type equations is from direct assignments to varibales or parameters: For the function call `bar(x)` to a function defined as `func bar(p type)` the resulting type equation would be `typeof(p) :≡ typeof(x)`. This is trivial to solve if there are no type parameters is either `p` or `x`. Further type equations can also be obtained from the constraints of type parameters: Given the function definition `func test[A Number]`, the type parameter `A` must satisfy the contraint (in this case being a Number). Therefore the type equation `A ∈ Number` can be added to the list.

Example of all type equations collected from a specific piece of code which makes use of type inference:

```
func checkEquality[E comparable](a, b E) bool {...}

func removeDuplicates[S ~[]P, P any](col S, eq func(P, P) bool) S

type Collection []string
var set Collection

filteredCollection := removeDuplicates(set, equalityCheck)
Explicit type arguments: einbauen
```

**Type parameteres + constraints**:
 - `S ~[]P`: The underlying type of S must be a slice of type p
 - `P any`: P can be of any type
 - `E comparable`: The type E must be comparable via `==` or `!=`

**Explicit type arguments**: none

These elements result in the following **type equations to be solved**:
- `S :≡ Collection`: The type `S` must be assignable to a `Collection`
- `func(P, P) bool :≡ func(E, E) bool`: The function type passed into `removeDuplicates` must be assignable to the type declared in the function
- `S ∈ ~[]P`: S must satisfy the constraint `~[]P` which means its an element of the type set of `~[]P`
- `P ∈ any`: P must satisfy the constraint `any` which is trivial
- `E ∈ comparable`: E must satisfy the constraint of being a `comparable` type

These type equations **result** in:
- S ➞ Collection
- E ➞ string
- P ➞ string


Another short example which also includes explicit type arguments:

```
type Printable interface{
    toString()
}

func printInput[E Printable](input E) {input.toString()}

type Rectangle struct{
    width int,
    height int
}

func (r Rectangle) toString(){...}

r := Rectangle{width: 2, height: 3}

printInput[Rectangle](r)
```
**Type parameteres + constraints**:
- `E Printable`: Type `E` must implement the interface `Printable`

**Explicit Type Arguments**:
- `E :≡ Rectangle`

These elements result in the following **type equations to be solved**:
- `E ∈ Printable`: The type E must be element of the type set of the constraint `Printable` which contains all types which implement the interface `Printable`
- `E :≡ Rectangle`: The assignment of Rectange to E is given explicitly in the Code

These type equations **result** in:
- E ➞ Rectangle

This example shows how the same mechanisms of sloving a set of type equations is also applied in cases where the types are explicitly statet and are therefore trivial to determine.


#### Solving these equations



Bound vs free type parameters...

