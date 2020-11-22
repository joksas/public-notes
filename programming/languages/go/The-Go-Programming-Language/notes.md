---
title: "The Go Programming Language"
---

These are my notes for the textbook "The Go Programming Language" by Alan Donovan and Brian Kernighan.

\clearpage

# Chapter 1

## Section 1.2

* Comments begin with `//`.
* Variables can be initialized as part of their declaration. If they are not explicitly initialized, they are implicitly initialized to *zero value*, i.e. `0` for numeric types and to an empty string for strings.
* `+` concatenates strings.
* `:=` is used for *short variable declaration*---it can declare one or more variables and give them appropriate types based on the initializer values.
* `i++` adds `1` to `i`. This is a statement, and not an expression, thus `j = i++` is illegal.

### `for` loops

#### Traditional

The `for` loop is *the only loop statement in Go*. One of its forms is
```go
for initialization; condition; post {
	// zero or more statements
}
```
For example:
```go
sum := 0
for i := 0; i < 10; i++ {
	sum += i
}
fmt.Println(sum)
```
* In the `for` loop, the opening brace must be on the same line as the `post` condition.
* `initialization` statement is executed before the loop starts.
* `condition` is a boolean expression evaluated at the beginning of each iteration.
* `post` statement is executed after the body of the loop, then the condition is evaluated again.
* The loop ends when the condition becomes false.
* Any of the `initialization`, `condition` or `post` can be omitted
* If neither of the last two are included, the semicolons can be omitted as well.

`while` loop can be constructed using `for` loop:
```go
// a traditional "while" loop
for condition {
	// ...
}
```

Similarly, and infinite loop can be constructed by omitting `condition` entirely:
```go
// a traditional infinite loop
for {
	// ...
}
```
Such a loop could still be terminated using `break` or `return` statement, for example.

#### `range`

A different form of the `for` loop can iterate over a *range* of values. For example:
```go
nums := []int{2, 0, 5, 2}
for idx, val := range nums {
	fmt.Println("index:", idx)
	fmt.Println("value", val)
}
fmt.Println(s)
```
* In each iteration, `range` produce two values: the index and the value of the element at that index.
* If one does not need an index, it can be assigned to *blank identifier*, `_`. This identifier can be used whenever syntax requires a variable name, but the program logic does not.
* If one does not need the value, it can be omitted (together with the comma)

### Declaring string variables

Here are four equivalent ways to declare a string variable:
```go
s := ""
var s string
var s = ""
var s string = ""
```
* The first one is the most compact but may only be used within a function, not for package-level variables.
* The second one relies on default initialization to an empty string.
* The third one is rarely used, except when declaring multiple variables.
* The fourth one is explicit about the variable's type, which is redundant. 

In practice, one should generally use either the first or the second form, with explicit initialization to say that the initial value is important and implicit initialization when it is not important. 

## Section 1.3

Functions and other package-level entities may be declared in any order.

### `map`

A `map` holds a set of key/value pairs. The key be of any type whose values can be compared with `==`. In the following example, function `make` creates a new empty map, in which the keys are strings and the values are `int`s:
```go
my_map := make(map[string]int)
```

One can iterate over the `map` using range:

```go
for key, val : = range my_map {
	// ...
}
```

The order of map iteration is not specified, but in practice it is random. This is intentional as it prevents programs from relying on any particular order where none is guaranteed.

`map`s are references to the data structures created by `make`. When a `map` is passed to a function, that function receives a copy of the reference. Any changes the function makes to the data structure will be visible through the caller function's `map` reference too.

### Scanners

One can create a scanner that reads from the program's standard input with the following command:
```go
input := bufio.NewScanner(os.Stdin)
```

With each call of `input.Scan()`, it reads the next line and removes the newline character from the end of the line. The result is retrieved by calling `input.Text()`. The `Scan` function return `true` when there is a line and `false` when there is not. 

Additionally, one can read files. With `os.Open(file)` one can open the file: 
```go
for _, arg := range files {
	f, err := os.Open(arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		continue
	}
	doThings(f)
	f.Close()
}
```

`f`, the first return value of `os.Open` is an open file that can be read by the `Scanner`. `err`, the second return value of `os.Open` is a value of the built-in error type. If `err == nil`, it means that the file was opened successfully. After the file is read, `Close` closes the file and releases any resources.

Finally, one can also read the entire input into memory at once and then manipulate it accordingly. This is done with the following piece of code:
```go
data, err := ioutil.ReadFile(filename)
```

### `Printf`

`fmt.Printf` allows to produce formatted output. Its first argument is a format string that specifies how the subsequent list of arguments should be formatted. The format of each argument is determined by a conversion character (a letter that follows a percent sign). For example, `%d` can format an integer using decimal notation. The following are some of the conversions:

| Conversion character |                       Displays argument as                      |
|:--------------------:|:---------------------------------------------------------------:|
| `%d`                 | Integer (base 10)                                               |
| `%b`                 | Integer (base 2)                                                |
| `%f`                 | Floating-point (no exponent)                                    |
| `%e`                 | Floating-point (with exponent)                                  |
| `%g`                 | Floating-point (`%e` for large exponents, `%f` otherwise)       |
| `%t`                 | Boolean                                                         |
| `%s`                 | String                                                          |
| `%c`                 | Character (represented by the corresponding Unicode code point) |
| `%q`                 | Quoted string or character                                      |
| `%v`                 | The value in a default (natural) format                         |
| `%%`                 | Literal percent sign                                            |

Additionally, there are *escape sequences* `\n` and `\t` for newline and tab, respectively. `Printf` does not write newline by default and so (by convention) don't any formatting functions whose names end in `f`, e.g. `log.Printf` or `fmt.Errorf`. Those functions whose names end in `ln` follow the logic of `Println`, i.e. formatting their arguments as if by `%v`, followed by `\n`.

## Section 1.4

### Importing packages

If we import a package whose path has multiple components, for example `image/color` and `image/gif`, in the following way
```go
import (
	"image"
	"image/color"
	"image/gif"
)
```
then we refer to the package with a name that comes from the last component. For example, `color.White` belongs to `image/color`, while `gif.GIF` belongs to `image/gif`.

### Constants

`const` declaration gives names to constants, values that are fixed at compile time. `const`s can be declared at package level (in which case their names would be visible throughout the package) or within a function. The value of a `const` must be a number, string or boolean. One can declare `const`s in the following way:
```go
const (
	myFirstConstant = 3.14
	mySecondConstant = 42
)
```

## Section 1.5

* Function `io.Copy(dst, src)` reads from `src` and writes to `dst`. One can use it for copying contents to `os.Stdout`, for example.
* Functions `strings.HasPrefix(s, prefix)` tests whether string `s` begins with `prefix`.


## Section 1.6

A *goroutine* is a concurrent function execution. This can be initiated in the following way:
```go
go myFunction(arg1, arg2) // start a goroutine
```
The function `main` runs in a goroutine as well.

A *channel* is a communication mechanism that allows one goroutine to pass values of *a specified type* to another goroutine.  To make a channel for strings:
```go
ch := make(chan string)
```
It can be received by a function:
```go
func myFunction(myChannel chan<- string) {
	myChannel <- "I am sending this message through myChannel to the caller function"
}
```
In the caller function, one can receive from the channel:
```go
fmt.Println(<-ch)
```

## Section 1.8

### Switches

Go also implements switches. For example, `switch` can operate on an operand:
```go
switch coinFlip() {
case "heads":
	heads++
case "tails":
	tails++
default:
	fmt.Println("Landed on edege.")
}
```
or it can simply list cases, each of which is a boolean expression (equivalent to `switch true`):
```go
func sumInt(x int, y int) {
	switch {
	case x + y < 0:
		fmt.Println("The sum is negative")
	default:
		fmt.Println("The sum is zero")
	case x + y > 0:
		fmt.Println("The sum is positive")
	}
}
```

`case`s are evaluated from top to bottom and the first matching is executed. The optional `default` is executed if *none* of the other cases match; it may be placed anywhere as seen in the example above. Cases do not fall through from one to the next one, though `fallthrough` statement can be used to override this behavior.

### Comments

As seen before, comments can be written with a pair of forward slashes (`//`). One can also use `/*` for the start of a block comment and `*/` for its end. However, for documenting code `//` are preferred.

# Chapter 2

## Section 2.1

In Go, the names of functions, variables, constants, types, statement labels and packages start with a letter or an underscore and may be followed by any number of additional letters, digits and underscores. The names can be set to anything, except for 25 keywords like `func` or `if`. Additionally, there are some predeclared names like `int` or `true` which can be redeclared (though that should be avoided in most cases).

If an entity is declared within a function, it is local to that function. If declared outside of a function, it is visible in all files of the package to which it belongs.

*The case of the first letter of name determines its visibility across package boundaries.* If a name begins with an uppercase letter, it is exported---it is visible outside of its own package and may be used by other parts of the program that imports that package, e.g. `fmt.Println()`. Package names are always in lowercase.

There is no limit on name length, but by convention shorter names are preferred in Go. Generally, the larger the scope of a name, the longer and more meaningful it should be.

In Go, camelCase is preferred, e.g. `MyFirstFunction()` and `mySecondFunction()` instead of `My_first_function()` and `my_second_function()`. The letters of acronyms and initialisms are always rendered in the same case, i.e. either only lowercase or only uppercase. For example, we would use `ufoGenerate()`, `UFOGenerate()`, `generateUFO()` or `GenerateUFO()`, but not `generateUfo()`.

## Section 2.2

A **declaration** names a program entity and specifies some or all of its properties. There are four major types of declarations: `var`, `const`, `type` and `func`.

A Go program is stored in or more files, all with extension `.go`. Each file begins with a package declaration that states which package the file is part of, e.g.:
```go
package main
```
After that follow `import` declarations and then package-level declarations of types, variables and functions in any order.

### Functions

A function declaration has a name, a list of parameters, an optional list of results and the function body. The result list is omitted if the function does not return anything. The function execution begins with the first statement and ends with the first encounter of a `return` statement or when the end of a function is reached that has no results. For example:
```go
func kmPerHToMPerS(kmPerH float64) float64 {
	return kmPerH * 1000 / 3600
}
```

## Section 2.3

A `var` declaration creates a variable of a particular type, attaches a name to it, and sets its initial value. In Go, there is no such thing as uninitialized variables. Go uses the zero-value mechanism to ensure that every variable holds a well-defined value of its type.

Each `var` declaration has the general form
```go
var name type = expression
```

Either `type` or `= expression` may be omitted, but not both because the variable has to be initialized:

* if `type` is omitted, the type is deduced from the initializer expression.
* if `= expression` is omitted, the initial value is set to the zero value of the specified type.

The table below lists zero values for various types.

| Type                                                       | Zero value                           |
|------------------------------------------------------------|--------------------------------------|
| All integer types                                          | `0`                                  |
| All float types                                            | `0.0`                                |
| Boolean                                                    | `false`                              |
| Strings                                                    | `""`                                 |
| Interfaces, slices, channels, maps, pointers and functions | `nil`                                |
| Array or struct                                            | Zero value for each element or field |

Go also allows to initialize more than one variable at a time. By omitting the type one can also declare variable with different types:
```go 
var a, b, c float64					// float64, float64, float64
var d, e, f = 5, 3.14, false		// int, float64, bool
```

A set of variables can also be initialized by calling a function that returns multiple values:
```go
var f, err = os.Open(name)	// a file and an error are returned
```

### Subsection 2.3.1

Inside a function, one may also use **short variable declaration** to declare and initialize variables. It has the following form:
```go
name := expression
```

The type of `name` is determined by the type of `expression`. Here is a couple of examples:
```go
anim := gif.GIF{LoopCount: nframes}
t := 0.0
```

Short variable declarations are used to declare most of local variables. `var` declaration is usually reserved for local variables that need an explicit type that differs from that of the initializer expression or when the variable will be assigned a value later and its initial value is not important.

As with `var` declarations, multiple variables can be declared using short variable declaration:
```go
i, j := 0, 1
```
but such declarations should be reserved only when they improve readability.

In Go, `:=` is declaration and `=` is assignment.

As with `var` declarations, short variable declaration may be used for calls to functions that return two or more values. However, short variable declaration does not necessarily declare all the variables on the LHS. If some of them were already declared in the same lexical block, then the short variable declaration acts like assignment to those variables. A short variable declaration must declare at least one new variable though.
```go
f, err := os.Open(infile)
// ...
out, err := os.Create(outfile) // no error because `out` newly declared
// ...
f, err := os.Create(outfile) // compile error: no new variables
```
One would have to use an ordinary assignment for the third statement.

### Subsection 2.3.2

A **variable** is a piece of storage containing a value. A **pointer** is the address of a variable, or simply the location where a value is stored.

If we declare a variable as
```go
var x int
```
then the expression `&x` yields a pointer to an integer variable. It is a value of type `*int`, which is pronounced as "pointer to `int`". If this pointer is called `p`, then we say "`p` points to `x`". The variable to which `p` points is written `*p`. It can even used in assignment and thus update the variable:
```go
x := 1
p := &x
fmt.Println(*p) // output: 1
*p = 2
fmt.Println(x) // output: 2
```

Each component of a variable of aggregate type (struct or array) is also a variable and thus has an address too.

The zero value of a pointer of any type is `nil`.

### Subsection 2.3.3

Built-in function `new` can be used to create variables as well. Specifically, expression
```go
new(type)
```
allows to create an unnamed variable of type `type`. This expression initializes the variable to the zero value of `type` and returns its address, which is a value of type `*type`.

The only difference with `new` is that one does not need to create a variable name, thus this feature is just a syntactic convenience.

Each call to `new` returns a distinct variable with a unique address. The only exception may be variables whose type carries no information and is thus of size zero, e.g. `struct{}` or `[0]int`.

## Section 2.4

The value of a variable is updated by an assignment statement. It has the following form:
```go
variable = expression
```
For example:
```go
x = 1
x = 2
x = x*5
```

Each of the arithmetic and bitwise binary operators has a corresponding **assignment operator**. For example:
```go
x *= 2 // equivalent to `x = x*2`
y -= 10 // equivalent to `y = y - 10`
```

Additionally, numeric variables can be incremented and decremented using `++` and `--` statements:
```go
x := 0
y := 0
x++ // `x` is now equal to `1`
y-- // `y` is now equal to `-1`
```

### Subsection 2.4.1

**Tuple assignment** allows several variables to be assigned values at once. All of the RHS expressions are evaluated before any of the variables are updated. This makes this feature useful when variables appear on both sides of the assignment; it can often eliminate the need to create temporary variables. For example, if we wanted to swap the values of two variables we could do
```go
x, y = y, x
```
instead of
```go
temp := x
x = y
y = temp
```

Some expressions, such as calls to functions with multiple results, produce several values; tuple assignments can be used there too.

### Subsection 2.4.2

In a lot of cases, assignments occur *implicitly*. For example:

* A function call implicitly assigns the argument values to the corresponding parameter variables
* A `return` statement implicitly assigns the values of the variables coming after it to the corresponding result variables
* A declaration or assignment for a composite type variable can result in implicit assignment of individual elements of that variable.

For example, declaration
```go
colors := []string{"red", "blue", "green", "yellow"}
```
implicitly assigns the following:
```go
colors[0] = "red"
colors[1] = "blue"
colors[2] = "green"
colors[3] = "yellow"
```

## Section 2.5

Type of a variable or expression characterizes the values it may take on. However, there are often variables that share the same representation but signify very different concepts. For example, `float64` could be used to represent both speed in kilometers per hour and speed in meters per second.

`type` declaration defines **named type** that has the same **underlying type** as an existing type. Named types allow to separate different and possibly incompatible uses of the underlying type. The declaration has the following structure:
```go
type name underlying_type
```

For example, we might use type declarations to differentiate between two units of speed, so that they could not be combined in expressions without a proper conversion:
```go
package speedconv

import "fmt"

type MPerS float64
type KmPerH float64

const (
	HumanWalkSpeed MPerS = 1
	HumanRunSpeed MPerS = 10
)

func MPSToKPH(speedMPS MPerS) KmPerH {
	return KmPerH(speedMPS * 3600 / 1000)
}

func KPHToMPS(speedKPH KmPerH) MPerS {
	return MPerS(speedKPH * 1000 / 3600)
}
```
here `KmPerH()` and `MPerS()` are conversions, not function calls. These do not change the value or representation in any way, just the named type.

For every type `T` there is a corresponding **conversion operation** `T(x)` that converts the value `x` to type `T`. Such conversion is allowed if both `x` and `T` have the same underlying type or if both are unnamed pointer types that point to variables of the same underlying type.

Conversions are also allowed between numeric types and between string and some slice types. However, these conversions may change the representation of the value. For example, converting floating point number to an integer will discard any fractional part.

The underlying type of a named type determines the intrinsic operations it supports. Comparison operators can be used to compare a value of a named type to another value of the same type, or to a value of an unnamed type with the same underlying type.

Named types also allow to define new behaviors for values of that type. These are expressed through functions associated with the type called type's **methods**. In these methods, parameter of the named type comes before the name of the method. For example:
```go
func (speedKPH KmPerH) String() string {
	return fmt.Sprintf("%g m/s", speedKPH)
}
```
`String()` method is often defined because it controls how values of the type appear when printed as a string by the `fmt` package.

## Section 2.6

Package initialization begins by firstly initializing package-level variables in the order in which they are declared, except when dependencies need to be resolved:
```go
var a = 5 // initialized first
var b = c + 1 // initialized third 
var c = 3 // initialized second 
```

In Go, one can use any number of `init` functions to perform some actions whenever the program starts. They are executed in the order in which they are declared:
```go
func init() {
	fmt.Println("Called first.")
}

func init() {
	fmt.Println("Called second.")
}
```

## Section 2.7

The **scope** of a declaration is the part of the source code where a use of the declared name refers to that *declaration*.

A **syntactic block** is a sequence of statements enclosed in *curly braces*, e.g. like in a body of a function of a loop. A name *declared* inside a syntactic block is not visible outside the block. More generally, we can define **lexical blocks** that are not necessarily surrounded by curly braces, but can still define scope. This expands the possible blocks to the **universe block**, **package block** and **file block**. Additionally there are lexical blocks for `for`, `if` and `switch` statements. One can also have simply curly braces without any special keyword.

The following code
```go
x := 0
fmt.Println(x)
{
    x = 1 // assignment changes the variable declared outside the block
    fmt.Println(x)
}
fmt.Println(x)
```
produces the output
```text
0
1
1
```

On the other hand, the following code
```go
x := 0
fmt.Println(x)
{
    x := 1 // a variable with the same name is declared inside the block 
    fmt.Println(x) // this will refer to the declaration
    // in the innermost enclosing lexical block
}
fmt.Println(x) // this will refer to the first declaration of `x` because it
// does not have access to names declared inside lexical blocks it is not a part of
```
produces the output
```text
0
1
0
```

Again, not all lexical blocks are denoted by explicit brace-delimited sequences, some are implied. For example, `for` loop creates two lexical blocks: the explicit block for the body, and an implicit block that is created by declaring variables using the initialization clause. The scope of a variable declared within this implicit block is the condition, post-statement and the body of the loop. For example
```go
x := "hello"
fmt.Println(x)
for _, x := range x { // refers to `x` declared outside the loop
	x := x + 'A' - 'a' // refers to `x` declared in the initialization clause 
	fmt.Printf("%c", x) // refers to `x` declared in the line above
}
fmt.Println()
fmt.Println(x) // refers to `x` declared outside the loop 
```
produces output
```text
hello
HELLO
hello
```

Similarly with `if` statements, the following code would produce three compile errors:
```go
if f, err := os.Open(fname); err != nil { // compile error: unused: f
	return err
}
f.Stat() // compile error: undefined f
f.Close() // compile error: undefined f
```
The way `f` and `err` are declared, they would only be accessible in the comparison statement and the body of `if`. One should instead rewrite it as
```go
f, err := os.Open(fname)
if err !nil {
	return err
}
f.Stat()
f.Close()
```
One might also try doing the following:
```go
if f, err := os.Open(fname); err != nil {
	return err
} else {
	// `f` and `err` are visible inside `else` too
	f.Stat()
	f.Close()
}
```
but that is not a recommended practice---successful execution path should not be indented.

Also, issues may occur with variables that are not used in the blocks they are declared, but should be used outside them. For example:
```go
var cwd string 

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("%v", err)
	}
}
```
will result in compile error because `cwd` is not used inside the `init()` function. One can avoid this error by declaring `err` in a separate `var` declaration, and then assigning values to `cwd` and `err`, instead of declaring them:
```go
var cwd string 

func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("%v", err)
	}
}
```

# Chapter 3

Go'data types are divided into four categories:

* basic types
* aggregate types
* reference types
* interface types

## Section 3.1

Go provides both signed and unsigned integer arithmetic. There are four distinct sizes of integers:

* 8 bits
* 16 bits
* 32 bits
* 64 bits

Signed integers of these sizes are implemented using types `int8`, `int16`, `int32` and `int64`, respectively. Similarly, unsigned versions are represented by types `uint8`, `uint16`, `uint32` and `uint32`, respectively.

Additionally, there are two types called `int` and `uint`. Both of them have the same size---either 32 or 64 bits---but one does not must not make any assumptions because different compilers may make different choices even on identical platforms. These two types are often the natural or most efficient size.

Type `rune` is a synonym for `int32`; it is usually used to indicate that a value is a Unicode code point. Two names can be used interchangeably. Similarly, type `byte` is a synonym for `uint8` and emphasizes that the value is a piece of raw data, and not a small numeric quantity.

Finally, there exists type `uintptr`, whose width is not specified but is sufficient to hold all the bits of a pointer value. It is usually used only for low-level programming.

Signed numbers are represented in 2's complement form and so the range of values that can be represented with an $n$-bit signed number is from $-2^{n-1}$ to $2^{n-1}-1$. Unsigned integers, on the other hand, use all bits to represent the value and so the range of values that can be represented with an $n$-bit unsigned number is from $0$ to $2^n-1$.

Go's binary operators for arithmetic, logic and comparison are listed below in the order of decreasing precedence:
```go
*    /    %    <<   >>   &    &^
+    -    |    ^
==   !=   <    <=   >    >=
&&
||
```

Operators at the same level associate to the left, so parentheses may be needed to override the order, e.g. `& (1 << 28)`.

Each operator in the first two lines in the table above has a corresponding **assignment operator**, e.g. `+=` for `+`.

`+`, `-`, `*`, and `/` can be applied to any integer, floating-point or complex number. However, remainder operator `%` can be applied only to integers. Its behavior for negative number varies across programming languages. In Go, the sign of the remainder is always the same as the sign of the dividend (numerator). If `/` is used with floating-point numbers, then the result is as expected, e.g. `10.0/4.0` is `2.5`. However, when used with integers, the result is truncated towards zero, e.g. `10/4` is `2`.

If the result of an operation has more bits than can be represented by the result type, it is said to **overflow**. The high-order bits that do not fit are silently discarded:
```go
var u uint8 = 255
fmt.Println(u, u+1) // "255 0"

var i int8 = 127
fmt.Println(i, i+1) // "127 -128"
```

Although unsigned numbers might seem natural for things like lengths of arrays, they should usually be avoided because they might easily result in unintended behavior. Suppose we wanted to print an array in reverse order:
```go
myList := []string{"first", "second", "third"}
for i := myLength(myList)-1; i >= 0; i-- {
	fmt.Println(myList[i])
}
```

If function `myLength()` returns a signed integer, the for loop will serve its function. However, if it returns an unsigned integer, the variable `i` will also be declared as unsigned integer. As a result, after `i` becomes `0` in one of the cycles, it would then change not to `-1` in the next cycle, but to $2^{64}-1$ (if the size of unsigned integer is 64 bits) instead. The loop would not terminate the way it was designed to. This is the reason why the built-in function `len()` in Go returns signed integers.

Conversion operations can be used to combine variables of different types:
```go
var apples int32 = 1
var oranges int64 = 2
var letsCompare int = int(apples) + int(oranges)
```

When applied to floating-point types, integer conversions will discard fractional part, truncating toward zero:
```go
myFloat := 2.9
myInt := int(myFloat)
fmt.Println(myFloat, myInt) // "2.9, 2"
```

Octal numbers can be written by adding `0` in front of them, e.g. `0512`. Similarly, hexadecimal numbers can be written by adding `0x` or `0X` in front of them, e.g. `0x3ba1` or `0Xf76aa1`. Nowadays, octal numbers seem to be used only for file permissions on POSIX systems.

Rune literals are written as a character within single quotes:
```go
asciiChar := 'a'
unicodeChar := 'ą'
```
## Section 3.2

Go uses two types of floating-point numbers, `float32` and `float64`. The former uses around 6 decimal digits of precision, while the latter around 15. Therefore, `float64` should be preferred in most cases because otherwise the errors can accumulate quickly.

Digits may be omitted before the decimal point (`.5`) or after it (`2.`). Number can also be written using scientific notation, e.g. `1.67e-27`. Additionally, there exist values of `+Inf` and `-Inf`, as well as `NaN`.

## Section 3.3

Go uses two types of complex numbers, `complex64` and `complex128`; they have real and imaginary components of types `float32` and `float64`, respectively. Function `complex()` creates a complex number from real and imaginary parts. Functions `real()` and `imag()` extract real and imaginary parts of a complex number, respectively. One can also simply define a compex number by adding an `i` after a decimal, for example:
```go
x := 4 + 3i
```

## Section 3.4

Type `bool` can take two values: `true` or `false`. `!` is logical negation, `&&` is logical AND and `||` is logical OR. Boolean values can be combined with several operators at a time, all of which have **shortcircuit behavior**: if the answer is already determined by left operand, the right operand is not evaluated. It makes it safe to write expressions like this, where `s` might be an empty string:
```go
if s != "" && s[0] == 'x'
```

## Section 3.5

In strings, index operation `s[i]` returns $i$-th byte of string `s`, e.g. `119`, while substring operation `s[i:j]` returns a new string from $i$-th to $j$-th byte, not including the latter, e.g. `hello w`.

A string value can be written inside double quotes. **Double-quoted string literals** can include escape sequences that begin with a backslash.

A **raw string literal** is enclosed in backticks; its contents are taken literally, including backslashes, thus no escape sequences are processed.

## Section 3.6

Constants are expressions whose value is known to the compiler; their evaluation occurs at compile time, and not at run time. The underlying type of any constant is a basic type: boolean, string or a number.

In Go, one usually does not specify a type for constants. Untyped constants can be used in expressions where they are combined with variables of certain type; if that constant can be converted to that type, no error will be thrown. Untyped constants also have much higher precision of 256 bits.

# Chapter 4

There are four composite types in Go: arrays, slices, maps and structs. Arrays and structs are fixed size, while slices and maps are dynamic data structures that can grow as more values are added. Arrays are homogeneous (all of their elements have the same type) while structs are heterogeneous (their elements can have different types).

## Section 4.1

Arrays can be declared in the following way:
```go
var a [3]int
```
This creates an array `a` of size 3.

We can also use **array literal** to initialize an array with specified values:
```go
var b [3]int = [3]int{1, 2, 3}
```

To remove redundancy of specifying the size of the array, we can use an ellipsis:
```go
c := [...]int{1, 2, 3}
```

Size of an array is part of its type, so, for example `[3]int` and `[4]int` are not equivalent.

We can also specify pairs of indices and values:
```go
type Temp float64

const (
	Celcius Temp = iota
	Farenheit
	Kelvin
)

symbol := [...]string{Celcius: "°C", Farenheit: "°F", Kelvin: "K"}
fmt.Println(Farenheit, symbol[Farenheit])
```

When the values are specified by the indices, these indices can appear in any order. If any are omitted, then they take the zero value of the type. For example,
```go
d := [...]int{99: 15}
```
defines array `d` of size 100. The first 99 elements are zero, and the last one has a value of `15`.

Array type is comparable if that array's element type is comparable. In that case, the comparison is performed by comparing individual elements.

In Go, if an array is passed as an argument to a function, it is received as a copy, i.e. it is not passed by reference. However, we can pass pointers to arrays and change them in this way. In Go, arrays are rarely used as function arguments or results; slices are usually used instead.

## Section 4.2

Slices are variable-length sequences. They are references to an **underlying array**. The **length** of a slice is the number of elements it contains. The **capacity** of a slice is the number of elements in the underlying array, counting from the first element in the slice.
```go
myArray := [6]string{"one", "two", "three", "four", "five", "six"}
myFirstSlice := myArray[1:4]
mySecondSlice := myArray[2:]
fmt.Println(len(myFirstSlice), cap(myFirstSlice)) // "3 5"
fmt.Println(len(mySecondSlice), cap(mySecondSlice)) // "4 4"
```

Slicing beyond capacity causes a panics, but slicing beyond length simply extends the slice.

Slices do not store the data themselves, they are just pointers to the underlying arrays. As a result, they can be used to modify the elements of those underlying arrays.

One can also use a **slice literal** to declare a slice:
```go
primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
```
This implicitly creates an array of size 8 and yields a slice `primes` that points to it.

Slices, unlike arrays, are not comparable with `==`; the comparison must be done manually.

Function `make` creates an array and returns a slice of it; the array is accessible only through the returned slice.
```go
make([]T, len)
make([]T, len, cap)
```

One can append a slice using function `append`. The first parameter of the function is the slice that needs to be appended. It is then followed by the values that will be appended. If the size of the underlying array (or, in other words, capacity of the slice) is too small, a new array is allocated. The returned slice points to that new array.

## Section 4.3

In Go, **maps** serve a very similar function as dictionaries do in Python. One of the key differences though is that in Go, values of a given map all have to have the same type.

A map type is written as `map[K]V` where `K` is the type of the key and `V` is the type of the value. Keys must be comparable with `==`, while values do not have restrictions on their type.

We can make a map and populate it in one of two ways. Using function `make`:
```go
wordcounts := make(map[string]int)
wordcounts["hello"] = 2
wordcounts["world"] = 1
```
or using a **map literal**:
```go
wordcounts := map[string]int{
  "hello": 2
  "world": 1
}
```

Consequently, an empty map can be declared using map literal as `map[string]int{}`.

One can access map elements using subscript notation, e.g. `wordcounts["hello"]`, and remove elements from it using `delete` function, e.g. `delete(wordcounts, "hello")`. If the key is not in the map, a zero value for its type is returned. To know if the element is non-existent, one can check the second return value when accessing an element. It is a boolean value that is `true` if the element is present:
```go
key = "hi"
count, ok := wordcounts[key]
if !ok {
  fmt.Printf("\"%s\" is not a key in map `wordcounts`!", key)
}
```

To cycle through all the elements of the map (which are not ordered), we can use `range`:
```go
for key, value := range myMap {
  fmt.Println(key, value)
}
```

## Section 4.4

A **struct** is an aggregate data type. Unlike arrays, structs can group values of arbitrary type. Each value is called a **field**.

We can declare a struct type in the following way:
```go
type Student struct {
  FirstName	string
  LastName	string
  ID		int
}
```
If a struct field begins with a capital letter, it is exported.

Once an instance of a struct type is declared, its fields can be accessed using dot notation:
```
var student1 Student
student1.FirstName = "John"
student1.LastName = "Doe"
student1.ID = 144
```

We can also use pointers:
```go
var bestStudent *Student = &student1
(*bestStudent).ID = 121
```

To avoid the cumbersome notation, Go allows to rewrite the last line as
```go
bestStudent.ID = 121
```

Values of a struct type can be specified using **struct literal** in two ways. It could be specified for every field, in the right order:
```go
type Position struct{
  X, Y int
}
myPosition := Position{4, 3}
```

This may not be optimal as one needs to remember the order of the fields. Additionally, it increases fragility because the code might change and so the definition of the struct type might be different later. A better approach is to list some or all of the fields and their values:
```go
myPosition := Position{
  X: 4
}
```

In Go, for efficiency larger struct types are often passed using a pointer. That is also the only way to modify the argument of a function.

If all the fields of a struct type are comparable, then two structs of that type can be compared using `==`.

Go also lets to declare struct fields with no name. These fields are called **anonymous fields**. For example:
```go
type Person struct {
  Name string
  Age int
}

type Employee struct {
  Person
  Salary int
}

var employee Employee
employee.Name = "Jane Doe" // equivalent to employee.Person.Name = "Jane Doe"
employee.Age = 50 // equivalent to employee.Person.Age = 50
employee.Salary = 65000
```

A **template** is a string that contains substrings enclosed in double curly brackets `{{...}}`; these are called **actions**. Inside actions, one can access variables and call functions. It provides an intuitive way of building things like HTML files (as evidenced by static website generator Hugo...).

# Chapter 5

## Section 5.1

Function declaration includes a **name**, a **list of parameters**, an optional **list of results**, and a **body**.

The results may be named or not. If they are named, they automatically initialize a variable to the zero value of its type. Thus the two declarations below are equivalent:
```go
func addA(x, y int) int {
  return x + y
}

func addB(x, y int) (z int) {
  z = x + y
  return
}
```

In Go, there is no concept of default parameter values, thus the arguments cannot be specified by name---they must be provided in the order in which the parameters were declared.

If a function declaration does not have a body, it indicates that the function is implemented in a language other than Go; it just specifies the function signature (parameters and returns).
