# Go Basics

![Go Gopher][gopher_image]

[gopher_image]: https://blog.golang.org/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg

This repository is designed to help beginners learn the fundamental concepts and syntax of the Go programming language. It contains a collection of examples and explanations covering various topics, including:

- Hello World
- Variables and Data Types
- Control Structures (if-else, switch, loops)
- Functions
- Arrays and Slices
- Maps
- Structs
- Interfaces
- Concurrency (Goroutines and Channels)
- Error Handling
- File I/O
- Packages and Modules

## Getting Started

### Installing Go

Before you can start working with this repository, you'll need to have Go installed on your machine. Follow these steps to install Go:

1. Download the latest version of Go from the official website: https://golang.org/dl/
2. Follow the installation instructions for your operating system.
3. After the installation is complete, verify that Go is installed by running the following command in your terminal or command prompt:

```go
go version
```
This should print the installed version of Go.

### Cloning the Repository

Once you have Go installed, you can clone this repository using the following command:

```
git clone https://github.com/your-username/go-basics.git
```

After cloning the repository, navigate to the respective folder for each topic to explore the examples and explanations.

## Usage

Each folder in this repository contains one or more Go source files (`.go` files) that demonstrate a specific concept or feature of the Go programming language. You can run these examples by navigating to the corresponding folder and executing the following command:

```go
go run hello/hello.go
```
This command will compile and run the `hello.go` file in the hello directory.

## Lexer

The lexer (short for lexical analyzer or scanner) is a component of the Go compiler or interpreter that breaks down the source code into a sequence of tokens. These tokens are the smallest meaningful units that the compiler or interpreter can understand, such as keywords, identifiers, operators, literals, and punctuation marks.

The lexer is the first stage of the compilation or interpretation process, and it is responsible for recognizing and classifying the input characters into meaningful tokens. It reads the source code character by character and applies a set of rules (defined by the language's lexical grammar) to identify and categorize the tokens.

In Go, the lexer is implemented in the `go/scanner` package, which provides a scanner that reads the source code and produces a stream of tokens. This repository includes examples and explanations on how to use the `go/scanner` package to tokenize Go source code.

When you write a program in a programming language like Go, you essentially write it in a human-readable form using words, symbols, and syntax rules defined by the language. However, computers can only understand machine-readable instructions, which are typically represented as binary code.
The lexer is the first step in the process of translating your human-readable code into a form that the computer can understand. Its job is to break down your source code into smaller, meaningful units called tokens.

Think of it like this: when you read a sentence in English, your brain automatically separates the sentence into individual words and punctuation marks. The lexer does a similar thing, but instead of separating words, it separates the source code into tokens like keywords (e.g., func, package), identifiers (e.g., variable and function names), operators (e.g., +, -, *), literals (e.g., numbers and strings), and punctuation marks (e.g., {, }, ;).

## Types in Go lang

Basic Types

* Boolean: bool (values are true or false)
* Numeric:

   * Integers: int8, int16, int32, int64, uint8, uint16, uint32, uint64, int (32 or 64 bits), uint (32 or 64 bits), uintptr (an unsigned integer large enough to store the uninterpreted bits of a pointer value)
   * Floating-point: float32, float64
   * Complex: complex64, complex128


* String: string (immutable sequence of bytes)
* Rune: rune (an alias for int32, represents a Unicode code point)

* Aggregate Types

    * Array: A numbered sequence of elements of a single type with a fixed length.
    * Slice: A reference to an underlying array with added functionality. Slices are dynamically-sized and can grow or shrink as needed.
    * Struct: A sequence of named elements called fields, each of which has a name and type.
    * Map: An unordered group of key-value pairs of the same type.
    * Channel: A channel is a communication mechanism used for sending and receiving values between goroutines (lightweight threads).

* Reference Types

    * Pointer: A pointer holds the memory address of a value of the same type.
    * Function: Functions are first-class values in Go and can be assigned to variables, passed as arguments, and returned from other functions.
    * Interface: An interface is a type that defines a set of method signatures. Any type that provides the methods defined by the interface can be considered an implementation of that interface.

## Variable Shadowing

Variable shadowing occurs when a variable declared in an inner scope has the same name as a variable declared in an outer scope, causing the inner variable to hide or "shadow" the outer variable within the inner scope.

## Comma ok or Error ok Syntax

In Go, the "comma ok" or "error ok" syntax is a way to check for successful operations and handle potential errors when using built-in functions or methods that return two values. The second value is typically a boolean or an error indicating whether the operation was successful or if an error occurred.
This is because in Go we don't have try catch block.

## Random Numbers

**math/rand** -  It provides a pseudorandom number generator (PRNG) for generating non-cryptographically secure random numbers, suitable for general-purpose applications.

**crypto/rand** - It provides a cryptographically secure pseudorandom number generator (CSPRNG) for generating unpredictable random numbers, suitable for security-sensitive applications like cryptography.

# Memory Management in Go

Go uses a tracing garbage collector to automatically manage memory. The garbage collector runs periodically in the background to free up unused memory. This means developers don't need to manually allocate or free memory. 

There are three main functions for memory allocation and initialization in Go:

- new() allocates memory for a type and returns a pointer. It is similar to malloc() in C. 

- make() allocates and initializes memory for maps, slices, and channels by preallocating the underlying data structure.

- delete() removes references to an object to allow it to be collected earlier by the garbage collector.

The garbage collector runs concurrently so it doesn't block the program during collection. This helps ensure responsiveness. Developers can control memory by removing references with delete().

## Pointers in Go

Pointers in Go allow us to reference a memory location that stores a value of a particular type. We use pointers because sometimes we need to pass large structures by reference instead of by value for efficiency.

In simple terms, a pointer contains the memory address of a value. We define a pointer using an asterisk (*) before the type. For example, a pointer to an int would be defined as *int. 

We can create a pointer by using the address of (&) operator on a variable, which gives us the memory location of that variable. Then we can pass that pointer to functions to modify the original variable.

Pointers are useful in Go because they let us modify values more efficiently for large data structures like slices, maps and channels that are passed by reference instead of copying by value. This avoids expensive copies of large amounts of data.

So in summary, pointers allow passing references to data instead of copies, improving performance for large or complex values.

## Arrays

Arrays in Go are fixed-length sequences of elements of the same type. Some key properties of arrays in Go include:

    - The length is part of the array's type, so [5]int and [10]int are different types. Arrays must be initialized to a fixed length.

    - You can declare and initialize an array like var a [5]int or a := [5]int{1, 2, 3, 4, 5}. Each element is initialized to the zero value by default. 

    - Unlike slices, arrays are values and can be directly assigned or passed to functions. However, their size is immutable so they cannot be appended to or resized like slices.

    - Common array operations include accessing elements via their index (a[0]), iterating with a for loop, and passing to functions to work with the elements.

## Slices

1. **Arrays vs. Slices**:
   - Arrays are like fixed-size boxes where you put things in. Once you define the size of an array, you can't change it.
   - Slices are like resizable boxes. They can grow or shrink as needed.

2. **Making a Slice**:
   - You can make a slice using the `make()` function or by slicing an existing array or slice.
   - Example: `s := make([]int, 5)` creates a slice of integers with a length of 5.

3. **Slicing**:
   - Slicing is like taking a piece of a pie. You specify a range of elements you want from an existing slice or array.
   - Example: `sliced := s[1:4]` creates a new slice containing elements from index 1 to index 3 of the original slice.

4. **Appending**:
   - You can add more items to a slice using the `append()` function. It makes the slice bigger.
   - Example: `s = append(s, 4)` adds the number 4 to the end of the slice.

5. **Length and Capacity**:
   - Every slice has a length, which is how many items it currently holds, and a capacity, which is how many items it can hold without resizing.
   - Example: `s := make([]int, 3, 5)` creates a slice with a length of 3 and a capacity of 5.

Slices are really handy because they let you work with collections of data easily, and they can change size to fit your needs.

# Maps in Go

In Go, a map is a built-in data structure that stores a collection of key-value pairs. It's like a dictionary or an associative array in other programming languages.

## Key-Value Pairs

- A map holds a collection of key-value pairs, where each key is unique within the map, and each key is associated with exactly one value.
- Keys and values can be of any data type, as long as the key types and value types are consistent within the map. For example, you can have a map where keys are strings and values are integers.

## Declaration and Initialization

```go
// Declaration and initialization using make()
var myMap map[string]int
myMap = make(map[string]int)

// Declaration and initialization using a map literal
anotherMap := map[string]int{"apple": 5, "banana": 3, "orange": 2}
```
## Accessing and Modifying Values

```go
// Accessing values
value := myMap["key"]

// Modifying values or adding new key-value pairs
myMap["newKey"] = 10

// Checking for key existence
value, ok := myMap["key"]
if ok {
    fmt.Println("Value exists:", value)
} else {
    fmt.Println("Key does not exist")
}

// Deleting key-value pair
delete(myMap, "key")
```

## Structs in Go

In Go, a struct is a composite data type that groups together zero or more fields of different data types under a single name. It's like a record or a structure in other programming languages.

## Definition

- You define a struct using the `type` keyword followed by the struct's name and a list of its fields enclosed in curly braces `{}`.
- Each field in the struct has a name and a data type.

```go
type Person struct {
    name string
    age  int
}
```

## Creating Instances

- You create instances (also known as objects or structs) of a struct by declaring variables of that struct type and initializing their fields.

```go
var p1 Person
p1.name = "Alice"
p1.age = 30

p2 := Person{name: "Bob", age: 25}
```

## Accessing Fields

- You can access fields of a struct using the dot (`.`) operator followed by the field name.

```go
fmt.Println(p1.name) // Output: Alice
fmt.Println(p2.age)  // Output: 25
```

## Anonymous Structs

- You can also create anonymous structs without specifying a name. These are useful for temporary data structures.

```go
p3 := struct {
    name string
    age  int
}{"Charlie", 35}
```

## Embedding

- You can embed one struct within another to create a hierarchy of data structures.

```go
type Address struct {
    city    string
    country string
}

type Person struct {
    name    string
    age     int
    address Address
}
```

Structs are commonly used in Go to represent complex data types, such as users, products, or any other entity that has multiple properties. They provide a way to organize related data and make the code more readable and maintainable.