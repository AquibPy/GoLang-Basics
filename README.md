# Go Basics

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