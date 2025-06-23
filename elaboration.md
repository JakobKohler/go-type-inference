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

As already mentioned above, the process of solving the type equations extracted above, both sides of the equation are compared against each other while recursively substituting types for their underlying types in order to find suitable type arguments that solve the equations. This process is called unification because the end goal is to unify both sides (i.e make the identical) so that moving on, they can be regarded as the same type. Two type parameters A and B are unified if they are matched against each other while neither of them already has an inferred type so now, if further down the type inferrence process either A or B are matched against a type T, they both are set to T at the same time.

**Examples illustrating the unification process:**

This example is going to show how the simple type equation `A ≡ B` can be solved for their respective type parameters. `A` and `B` are defined as follows: 
`A: struct{a map[E]F; b []int}`
`B: struct{a map[string]byte; b []G}`

In this example there are three type parameters which need to be determined: `E, F, G`. The algorithm starts the comparison at the root of the type tree, so in this case with `struct{...} ≡ struct{...}` So far, the compared types are identical with both sides being a function type. In this case, the unification algorithm descends recursively on both sides. It start by comparing the types of the first fields in the two struct types which results in the equation `map[E]F ≡ map[string]byte`. The procedure stays the same here due to the recursive nature of the unification and the new comparison becomes `map[...] ≡ map[...]` which is identical hence the process goes one layer deeper where it first compares the types of the map keys and then the types of the map values: So the keys result in the comparison `E ≡ string` from which it can instantly be inferred that the type of E must be string: `E ➞ string`. After that, the map value types are compared: `F ≡ byte => F ➞ byte`. That is the first struct field resolved. For the second field the resulting comparison is `[]int ≡ []G`. In this case, the algorithm also descends until there are no types in the equation anymore which have underlying types. By doing this, `[]int ≡ []G` results in `int ≡ G` which can also be solved trivially with `G ➞ int`. In summary, the recursive unification process has successfully inferred (or solved for) the types of all type parameters in the given equation.
```
E ➞ string
F ➞ byte
G ➞ int
```
In this process, mainly two types of issues can occur: Trying to unify types of different underlying structures and trying to unify types for which the type arguments create logically conflicting necessities.

The first issue is rather simple to detect and avoid. If the two types compared are not of the same underlying structure, the unification and hence the type inference fails. Example: `A: map[int]int, B: map[E][5]int`. The first comparision (`map[...] ≡ map[...]`) succeeds for both types (both are maps), in the second comparison one layer deeper, the algorithm correctly infers `E ➞ int` from `E ≡ int`. It fails with the final comparison of `int ≡ [5]int`.

Conflicting type arguments can sometimes be harder to spot but will also cause unification (and type inference) to fail. For this example, the types A and B are defined as follows: `A: struct{a E; b byte; c []E}` and `B: struct{a bool; b F, c []F}`. Recursivce type unification determines the following parinings correctly: `E ≡ bool => E ➞ bool` from the first field of the struct, `byte ≡ F => F ➞ byte`. The conflict arises with the unification of the third struct field which results in `[]E ≡ []F => E ≡ F`. This cannot be true if E and F are of different types so the unification fails at this point.

**Special situation: Untyped constants**

The above looked at only situations in which there are types on both sides of the equation. But it is quite common in code to find a situation in which there is no type specifically given. Untyped constants are usually not considered for type inference and is subordinate to typed arguments. Only in the situation where the type parameter has no inferred type yet, the untyped constant supplies this type. In Go, untyped constants have a default type which will be used for type inference in these cases. If different constants compete for the same type variable it usually results in an error unless the constants are in some order assignable to one of the competeing default type (e.g int is assignable to float64)
**Example**:

```
func test[P any](a P, b P) {...}
var a bool
```

`foo(a, a)`: Explicit type given => `P -> bool`
`foo(a, 3)`: Explicit type given => `P -> bool`, but second parameter `b` cannot be assigned to bool, so this creates an error
`foo(3, 4)`: Both parameters have the same default type (int) so: `P -> int`
`foo(3, 4.2)`: Default types `int` and `float64`. In this case Go selects the larger type to go with, so: `P -> float64`
`foo(3, "Test")`: Default types `int` and `string`





<!-- Bound vs free type parameters...
Untyped constants
Whats missing:
No simplifying, No X :≡ Y or X ∈ Y type relations -->

Quelle: Go Blog

